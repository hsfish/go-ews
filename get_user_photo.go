package ews

import (
	"encoding/xml"
	"errors"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type GetUserPhotoRequest struct {
	XMLName       struct{} `xml:"m:GetUserPhoto"`
	Email         string   `xml:"m:Email"`
	SizeRequested string   `xml:"m:SizeRequested"`
}

type GetUserPhotoResponse struct {
	ewsxml.Response
	HasChanged  bool   `xml:"HasChanged"`
	PictureData string `xml:"PictureData"`
}

type getUserPhotoResponseEnvelop struct {
	XMLName struct{}                 `xml:"Envelope"`
	Body    getUserPhotoResponseBody `xml:"Body"`
}
type getUserPhotoResponseBody struct {
	GetUserPhotoResponse GetUserPhotoResponse `xml:"GetUserPhotoResponse"`
}

// GetUserPhoto
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserphoto-operation
func GetUserPhoto(c Requester, r *GetUserPhotoRequest) (*GetUserPhotoResponse, error) {
	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.Request(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getUserPhotoResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.GetUserPhotoResponse.ResponseClass == ewsxml.ResponseClass_Error {
		return nil, errors.New(soapResp.Body.GetUserPhotoResponse.MessageText)
	}

	return &soapResp.Body.GetUserPhotoResponse, nil
}
