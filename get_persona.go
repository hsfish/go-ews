package ews

import (
	"encoding/xml"
	"errors"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type GetPersonaRequest struct {
	XMLName   struct{}         `xml:"m:GetPersona"`
	PersonaId ewsxml.PersonaId `xml:"m:PersonaId"`
}

type getPersonaResponseEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    getPersonaResponseBody `xml:"Body"`
}
type getPersonaResponseBody struct {
	FindPeopleResponse GetPersonaResponse `xml:"GetPersonaResponseMessage"`
}

type GetPersonaResponse struct {
	ewsxml.Response
	Persona ewsxml.Persona `xml:"Persona"`
}

// GetPersona
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getpersona-operation
func GetPersona(c Requester, r *GetPersonaRequest) (*GetPersonaResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.Request(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getPersonaResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.FindPeopleResponse.ResponseClass == ewsxml.ResponseClass_Error {
		return nil, errors.New(soapResp.Body.FindPeopleResponse.MessageText)
	}

	return &soapResp.Body.FindPeopleResponse, nil
}
