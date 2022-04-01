package ews

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/go-xmlfmt/xmlfmt"
)

type Logger interface {
	Server(url string, ver Version)
	DumpRequest(req *http.Request, body []byte)
	DumpResponse(resp *http.Response, body []byte)
}

func DefaultLogger() Logger { return new(logger) }
func NopLogger() Logger     { return new(nopLogger) }

type logger struct{}

func (l *logger) Server(url string, ver Version) {
	log.Println("EWS Server:", url, ver)
}

func (l *logger) DumpRequest(req *http.Request, body []byte) {
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		log.Println("Dump error:", err)
	}
	log.Printf(
		"Request:\n%s%s\n----\n",
		dump,
		xmlfmt.FormatXML(string(body), "", "  "),
	)
}

func (l *logger) DumpResponse(resp *http.Response, body []byte) {
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		log.Println("Dump error:", err)
	}
	log.Printf(
		"Response:\n%s%s\n----\n",
		dump,
		xmlfmt.FormatXML(string(body), "", "  "),
	)
}

type nopLogger struct{}

func (_ *nopLogger) Server(_ string, _ Version)              {}
func (_ *nopLogger) DumpRequest(_ *http.Request, _ []byte)   {}
func (_ *nopLogger) DumpResponse(_ *http.Response, _ []byte) {}
