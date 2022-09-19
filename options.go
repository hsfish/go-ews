package ews

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/Azure/go-ntlmssp"
)

type Option func(c *client) error

func WithLogger(l Logger) Option {
	return func(c *client) error {
		if l == nil {
			c.log = NopLogger()
		} else {
			c.log = l
		}
		return nil
	}
}

func WithDefaultLogger() Option { return WithLogger(DefaultLogger()) }

func WithBasicAuth(user, pass string) Option {
	return func(c *client) error {
		c.auth = [2]string{user, pass}
		return nil
	}
}

// WithTimeout sets the Timeout field on the http.Client.
func WithTimeout(t time.Duration) Option {
	return func(c *client) error {
		c.http.Timeout = t
		return nil
	}
}

func WithTransport(t http.RoundTripper) Option {
	return func(c *client) error {
		c.http.Transport = t
		return nil
	}
}

func WithDefaultTransport(skipTls bool) Option {
	t := http.DefaultTransport
	if skipTls {
		t.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return WithTransport(t)
}

func WithNTLM(skipTls bool) Option {
	return WithTransport(new(ntlmssp.Negotiator))
}

func WithExchangeImpersonation(v string) Option {
	return func(c *client) error {
		c.header.ImpersonateSmtpAddress(v)
		return nil
	}
}
