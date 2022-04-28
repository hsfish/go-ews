package ews

import (
	"context"
	"encoding/xml"
	"io"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
)

type Operation interface {
	Header() *ewsxml.Header
	Body() interface{}
}

type Request struct {
	ctx  context.Context
	head *ewsxml.Header
	body interface{}
}

func NewRequest(head *ewsxml.Header, body interface{}) *Request {
	return NewRequestWithContext(nil, head, body)
}

func NewRequestWithContext(ctx context.Context, head *ewsxml.Header, body interface{}) *Request {
	return &Request{
		ctx:  ctx,
		head: head,
		body: body,
	}
}

func NewOperationRequest(ctx context.Context, op Operation) *Request {
	return NewRequestWithContext(ctx, op.Header(), op.Body())
}

func (r *Request) Header() *ewsxml.Header { return r.head }
func (r *Request) Body() interface{}      { return r.body }

func (r *Request) Context() context.Context {
	if r.ctx == nil {
		r.ctx = context.Background()
	}
	return r.ctx
}

//goland:noinspection HttpUrlsUsage
var (
	soapStart = []byte(xml.Header + `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages"
		xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"
		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">`)

	soapBodyStart = []byte(`<soap:Body>`)
	soapEnd       = []byte(`</soap:Body></soap:Envelope>`)
)

func (r *Request) WriteBody(w io.Writer) error {
	x, err := xml.Marshal(r.head)
	if err != nil {
		return errors.WithStack(err)
	}

	w.Write(soapStart)
	w.Write(x)
	w.Write(soapBodyStart)

	switch b := r.body.(type) {
	case []byte:
		w.Write(b)

	case string:
		io.WriteString(w, b)

	default:
		if x, err = xml.Marshal(b); err != nil {
			return errors.WithStack(err)
		} else {
			w.Write(x)
		}
	}

	w.Write(soapEnd)
	return nil
}
