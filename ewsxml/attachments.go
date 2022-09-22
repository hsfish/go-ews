package ewsxml

import (
	"encoding/xml"
)

// The Attachments element contains the items or files that are attached
// to an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachments-ex15websvcsotherref
type Attachments struct {
	XMLName        xml.Name `xml:"Attachments"`
	ItemAttachment []ItemAttachment
	FileAttachment []FileAttachment
}

// The ItemAttachment element represents an Exchange item that is attached
// to another Exchange item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemattachment
type ItemAttachment struct {
	XMLName          xml.Name `xml:"ItemAttachment"`
	AttachmentId     AttachmentId
	Name             string
	ContentType      string
	ContentId        string
	ContentLocation  string
	Size             uint
	LastModifiedTime string
	IsInline         bool
	CalendarItem     CalendarItem
}

// The FileAttachment element represents a file that is attached to
// an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileattachment
type FileAttachment struct {
	XMLName          xml.Name `xml:"FileAttachment"`
	AttachmentId     AttachmentId
	Name             string
	ContentType      string
	ContentId        string
	ContentLocation  string
	Size             uint
	LastModifiedTime string
	IsInline         bool
	IsContactPhoto   bool
	Content          string
}

// The AttachmentId element identifies an item or file attachment.
// This element is used in CreateAttachment responses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentid
type AttachmentId struct {
	XMLName          xml.Name `xml:"AttachmentId"`
	Id               string   `xml:"Id,attr"`
	RootItemId       string   `xml:"RootItemId,attr,omitempty"`
	RootItemChangeId string   `xml:"RootItemChangeId,attr,omitempty"`
}

type SendAttachmentId struct {
	XMLName xml.Name `xml:"t:AttachmentId"`
	Id      string   `xml:"Id,attr"`
}

// The GetAttachment element is the root element in a request to get an attachment
// from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachment
type GetAttachment struct {
	XMLName         xml.Name `xml:"m:GetAttachment"`
	AttachmentShape AttachmentShape
	AttachmentIds   *AttachmentIds
}

type AttachmentIds struct {
	XMLName      xml.Name `xml:"m:AttachmentIds"`
	AttachmentId []SendAttachmentId
}

// The AttachmentShape element identifies additional properties to return
// in a response to a GetAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentshape
type AttachmentShape struct {
	XMLName            xml.Name `xml:"m:AttachmentShape"`
	IncludeMimeContent bool     `xml:"t:IncludeMimeContent,omitempty"`
	BodyType           BodyType `xml:"t:BodyType,omitempty"`
	FilterHtmlContent  bool     `xml:"t:FilterHtmlContent,omitempty"`
	// AdditionalProperties
}

// The GetAttachmentResponseMessage element contains the status and result of
// a single GetAttachment operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachmentresponsemessage
type GetAttachmentResponseMessage struct {
	XMLName     xml.Name `xml:"GetAttachmentResponseMessage"`
	Attachments Attachments
}
