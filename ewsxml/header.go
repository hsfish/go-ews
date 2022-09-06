package ewsxml

import (
	"encoding/xml"
)

type Version string

func (v Version) String() string { return string(v) }

type Header struct {
	XMLName               xml.Name               `xml:"soap:Header"`
	RequestServerVersion  RequestServerVersion   `xml:"t:RequestServerVersion"`
	ExchangeImpersonation *ExchangeImpersonation `xml:"t:ExchangeImpersonation,omitempty"`
	TimeZoneContext       *TimeZoneContext       `xml:",omitempty"`
}

func (h *Header) ServerVersion(ver Version) {
	h.RequestServerVersion.Version = ver
}

func (h *Header) DiscardImpersonation() { h.ExchangeImpersonation = nil }

func (h *Header) ImpersonateSmtpAddress(v string) {
	h.ExchangeImpersonation = &ExchangeImpersonation{
		ConnectingSID: ConnectingSID{SmtpAddress: v},
	}
}

func (h *Header) ImpersonatePrimarySmtpAddress(v string) {
	h.ExchangeImpersonation = &ExchangeImpersonation{
		ConnectingSID: ConnectingSID{PrimarySmtpAddress: v},
	}
}

type RequestServerVersion struct {
	Version Version `xml:"Version,attr"`
}

type ExchangeImpersonation struct {
	ConnectingSID ConnectingSID `xml:"t:ConnectingSID,omitempty"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectingsid
type ConnectingSID struct {
	PrincipalName      string `xml:"t:PrincipalName,omitempty"`
	SID                string `xml:"t:SID,omitempty"`
	SmtpAddress        string `xml:"t:SmtpAddress,omitempty"`
	PrimarySmtpAddress string `xml:"t:PrimarySmtpAddress,omitempty"`
}
