package ewsxml

import (
	"encoding/xml"
)

type ResponseClass string

func (r ResponseClass) String() string { return string(r) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// ResponseClass_Success describes a request that is fulfilled.
	ResponseClass_Success ResponseClass = "Success"
	// ResponseClass_Warning describes a request that was not processed. A
	// warning may be returned if an error occurred while an item in the
	// request was processing and subsequent items could not be processed.
	ResponseClass_Warning ResponseClass = "Warning"
	// ResponseClass_Error describes a request that cannot be fulfilled.
	// Information about the error can be found in the Response.ResponseCode
	// and Response.MessageText elements.
	ResponseClass_Error ResponseClass = "Error"
)

type ResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Response []byte `xml:",innerxml"`
	} `xml:"Body"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessage
type Response struct {
	ResponseClass ResponseClass `xml:"ResponseClass,attr"`
	// The MessageText element provides a text description of the status of the
	// response.
	// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetext
	MessageText  string       `xml:"MessageText"`
	ResponseCode ResponseCode `xml:"ResponseCode"`
	MessageXml   MessageXml   `xml:"MessageXml"`
}

func (r Response) String() string { return r.MessageText }

type MessageXml struct {
	ExceptionType       string `xml:"ExceptionType"`
	ExceptionCode       string `xml:"ExceptionCode"`
	ExceptionServerName string `xml:"ExceptionServerName"`
	ExceptionMessage    string `xml:"ExceptionMessage"`
}
