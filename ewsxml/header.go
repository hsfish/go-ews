package ewsxml

import (
	"encoding/xml"
)

type Version string

func (v Version) String() string { return string(v) }

type Header struct {
	XMLName               xml.Name `xml:"soap:Header"`
	RequestServerVersion  RequestServerVersion
	ExchangeImpersonation ExchangeImpersonation
}

func (h *Header) SetVersion(ver Version) {
	h.RequestServerVersion.Version = ver
}

func (h *Header) DiscardImpersonation() {
	h.ExchangeImpersonation = ExchangeImpersonation{}
}

func (h *Header) ImpersonateSmtpAddress(v string) {
	h.ExchangeImpersonation.ConnectingSID.SmtpAddress = v
}

func (h *Header) ImpersonatePrimarySmtpAddress(v string) {
	h.ExchangeImpersonation.ConnectingSID.PrimarySmtpAddress = v
}

type RequestServerVersion struct {
	XMLName xml.Name `xml:"t:RequestServerVersion"`
	Version Version  `xml:"Version,attr"`
}

type ExchangeImpersonation struct {
	XMLName       xml.Name      `xml:"t:ExchangeImpersonation"`
	ConnectingSID ConnectingSID `xml:"t:ConnectingSID,omitempty"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectingsid
type ConnectingSID struct {
	PrincipalName      string `xml:"t:PrincipalName,omitempty"`
	SID                string `xml:"t:SID,omitempty"`
	SmtpAddress        string `xml:"t:SmtpAddress,omitempty"`
	PrimarySmtpAddress string `xml:"t:PrimarySmtpAddress,omitempty"`
}
