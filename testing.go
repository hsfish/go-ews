package ews

import (
	"net/http"
)

func newTestClient(ver Version, t *testTransport) (Client, error) {
	return NewClient("test.client", ver, WithTransport(t))
}

type testTransport struct {
	req *http.Request
}

// RoundTrip catches the request and returns nil for both response and error.
func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.req = req
	return nil, nil
}

func (t *testTransport) GetRequest() *http.Request { return t.req }
