/*
Package ews (Exchange Web Service)

https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/ews-operations-in-exchange
*/
package ews

import (
	"bytes"
	"context"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
)

type Version = ewsxml.Version

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	Exchange2010     Version = "Exchange2010"
	Exchange2013_SP1 Version = "Exchange2013_SP1"

	RequestError   errors.Kind = "request error"
	UnmarshalError errors.Kind = "unmarshal error"
)

type Requester interface {
	// Request
	// Argument body must be of []byte or any type that xml.Marshal
	// successfully can handle.
	Request(ctx context.Context, body interface{}) ([]byte, error)
}

type Client interface {
	Requester
	Log() Logger
	Url() string
	Username() string
	Header() *ewsxml.Header
}

func NewClient(url string, ver Version, opts ...Option) (Client, error) {
	c := &client{
		log: NopLogger(),
		url: url,
		http: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Second * 10,
		},
	}

	var err error
	for _, opt := range opts {
		errors.Append(&err, opt(c))
	}

	c.header.SetVersion(ver)
	c.log.Server(url, ver)

	return c, err
}

type client struct {
	log    Logger
	http   *http.Client
	url    string
	auth   [2]string
	header ewsxml.Header
}

func (c *client) Log() Logger { return c.log }

func (c *client) Url() string { return c.url }

func (c *client) Username() string { return c.auth[0] }

func (c *client) Header() *ewsxml.Header { return &c.header }

func requestAndUnmarshal(ctx context.Context, req Requester, body interface{}, dest interface{}) error {
	data, err := req.Request(ctx, body)
	if err != nil {
		return err
	}
	return errors.WithKind(xml.Unmarshal(data, dest), UnmarshalError)
}

func (c *client) Request(ctx context.Context, body interface{}) ([]byte, error) {
	if ctx == nil {
		// default to context.Background() just like http.NewRequest()
		ctx = context.Background()
	}

	req, err := c.createRequest(ctx, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, errors.WithKind(err, RequestError)
	}
	defer errors.AppendFunc(&err, resp.Body.Close)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(resp)
	}

	var x ewsxml.ResponseEnvelope
	if err = xml.Unmarshal(data, &x); err != nil {
		return data, errors.WithKind(err, UnmarshalError)
	}

	return x.Body.Response, err
}

//goland:noinspection HttpUrlsUsage
const (
	soapStart = xml.Header + `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages"
		xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"
		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">`

	soapBodyStart = `<soap:Body>`
	soapEnd       = `</soap:Body></soap:Envelope>`
)

func (c *client) createRequest(ctx context.Context, body interface{}) (*http.Request, error) {
	buf, err := xml.Marshal(c.header)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf = append([]byte(soapStart), buf...)
	buf = append(buf, soapBodyStart...)

	switch b := body.(type) {
	case []byte:
		buf = append(buf, b...)

	case string:
		buf = append(buf, b...)

	default:
		if x, err := xml.Marshal(b); err != nil {
			return nil, errors.WithStack(err)
		} else {
			buf = append(buf, x...)
		}
	}

	buf = append(buf, soapEnd...)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Url(), bytes.NewReader(buf))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if c.auth[0] != "" {
		req.SetBasicAuth(c.auth[0], c.auth[1])
	}
	req.Header.Set("Content-Type", "text/xml")

	c.log.DumpRequest(req, buf)
	return req, nil
}
