/*
Package ews (Exchange Web Service)

https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/ews-operations-in-exchange
*/
package ews

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
)

type Version = ewsxml.Version

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	Exchange2010     Version = "Exchange2010"
	Exchange2013_SP1 Version = "Exchange2013_SP1"
)

type Requester interface {
	// Request
	// Argument body must be of []byte or any type that xml.Marshal
	// successfully can handle.
	Request(body interface{}) ([]byte, error)
}

type Client interface {
	Requester
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
	log       Logger
	http      *http.Client
	url       string
	basicAuth [2]string
	dev       bool

	header ewsxml.Header
}

func (c *client) Url() string { return c.url }

func (c *client) Username() string { return c.basicAuth[0] }

func (c *client) Header() *ewsxml.Header { return &c.header }

func requestAndUnmarshal(client Requester, body interface{}, dest interface{}) error {
	data, err := client.Request(body)
	if err != nil {
		return err
	}
	return errors.WithStack(xml.Unmarshal(data, dest))
}

func (c *client) Request(body interface{}) ([]byte, error) {
	req, err := c.createRequest(body)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer errors.AppendFunc(&err, resp.Body.Close)

	data, err := ioutil.ReadAll(resp.Body)
	c.log.DumpResponse(resp, data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(resp)
	}

	var x ewsxml.ResponseEnvelope
	if err = xml.Unmarshal(data, &x); err != nil {
		return data, errors.WithStack(err)
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

func (c *client) createRequest(body interface{}) (*http.Request, error) {
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

	req, err := http.NewRequest("POST", c.Url(), bytes.NewReader(buf))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if c.basicAuth[0] != "" {
		req.SetBasicAuth(c.basicAuth[0], c.basicAuth[1])
	}
	req.Header.Set("Content-Type", "text/xml")

	c.log.DumpRequest(req, buf)
	return req, nil
}
