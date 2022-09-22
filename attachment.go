package ews

import (
	"context"
	"encoding/xml"

	"github.com/soft-stech/go-ews/ewsxml"
)

type GetAttachmentOperation struct {
	header        ewsxml.Header
	GetAttachment ewsxml.GetAttachment
}

func (op *GetAttachmentOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetAttachmentOperation) Body() interface{}      { return op.GetAttachment }

type GetAttachmentResponse struct {
	XMLName          xml.Name `xml:"GetAttachmentResponse"`
	ResponseMessages struct {
		XMLName                      xml.Name `xml:"ResponseMessages"`
		GetAttachmentResponseMessage ewsxml.GetAttachmentResponseMessage
	}
}

func GetAttachments(ctx context.Context, req Requester, op *GetAttachmentOperation) (*GetAttachmentResponse, error) {
	var out GetAttachmentResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
