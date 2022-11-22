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

type CreateAttachmentOperation struct {
	header           ewsxml.Header
	CreateAttachment struct {
		ewsxml.CreateAttachment
	}
}

func (op *CreateAttachmentOperation) Header() *ewsxml.Header { return &op.header }
func (op *CreateAttachmentOperation) Body() interface{}      { return op.CreateAttachment }

type CreateAttachmentResponse struct {
	XMLName          xml.Name `xml:"CreateAttachmentResponse"`
	ResponseMessages struct {
		XMLName                          xml.Name `xml:"ResponseMessages"`
		CreateAttachmentsResponseMessage ewsxml.CreateAttachmentsResponseMessage
	}
}

type DeleteAttachmentOperation struct {
	header           ewsxml.Header
	DeleteAttachment ewsxml.DeleteAttachment
}

func (op *DeleteAttachmentOperation) Header() *ewsxml.Header { return &op.header }
func (op *DeleteAttachmentOperation) Body() interface{}      { return op.DeleteAttachment }

type DeleteAttachmentResponse struct {
	XMLName          xml.Name `xml:"DeleteAttachmentResponse"`
	ResponseMessages struct {
		XMLName                      xml.Name `xml:"ResponseMessages"`
		GetAttachmentResponseMessage ewsxml.DeleteAttachmentResponseMessage
	}
}

func GetAttachments(ctx context.Context, req Requester, op *GetAttachmentOperation) (*GetAttachmentResponse, error) {
	var out GetAttachmentResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func CreateAttachment(ctx context.Context, req Requester, op *CreateAttachmentOperation) (*CreateAttachmentResponse, error) {
	var out CreateAttachmentResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func DeleteAttachment(ctx context.Context, req Requester, op *DeleteAttachmentOperation) (*DeleteAttachmentResponse, error) {
	var out DeleteAttachmentResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
