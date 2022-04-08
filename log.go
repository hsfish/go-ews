package ews

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type Logger interface {
	Server(url string, ver Version)
	HttpRequest(req *http.Request, body []byte)
	HttpResponse(resp *http.Response, body []byte)
	Response(resp ewsxml.Response)
}

func DefaultLogger() Logger { return new(logger) }
func NopLogger() Logger     { return new(nopLogger) }

type logger struct{}

func (l *logger) Server(url string, ver Version) {
	log.Println("EWS Server:", url, ver)
}

func (l *logger) HttpRequest(req *http.Request, body []byte) {
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		log.Println("Dump error:", err)
	} else {
		log.Printf("Request:\n%s%s\n----\n", dump, body)
	}
}

func (l *logger) HttpResponse(resp *http.Response, body []byte) {
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		log.Println("Dump error:", err)
	} else {
		log.Printf("Response:\n%s%s\n----\n", dump, body)
	}
}

func (l *logger) Response(resp ewsxml.Response) {
	if resp.ResponseCode == ewsxml.NoError {
		return
	}
	log.Printf("%s: %s (%s)", resp.ResponseClass, resp.MessageText, resp.ResponseCode)
}

type nopLogger struct{}

func (_ *nopLogger) Server(_ string, _ Version)              {}
func (_ *nopLogger) HttpRequest(_ *http.Request, _ []byte)   {}
func (_ *nopLogger) HttpResponse(_ *http.Response, _ []byte) {}
func (_ *nopLogger) Response(_ ewsxml.Response)              {}
