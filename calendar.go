package ews

import (
	"context"
	"encoding/xml"

	"github.com/soft-stech/go-ews/ewsxml"
)

type FindItemCalendarViewOperation struct {
	header   ewsxml.Header
	FindItem struct {
		ewsxml.FindItem
		CalendarView ewsxml.CalendarView

		// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderids
		ParentFolderIds struct {
			DistinguishedFolderId ewsxml.DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
		} `xml:"m:ParentFolderIds"`
	}
}

func (op *FindItemCalendarViewOperation) Header() *ewsxml.Header { return &op.header }
func (op *FindItemCalendarViewOperation) Body() interface{}      { return op.FindItem }

type FindItemCalendarViewResponse struct {
	XMLName          xml.Name `xml:"FindItemResponse"`
	ResponseMessages struct {
		XMLName                 xml.Name `xml:"ResponseMessages"`
		FindItemResponseMessage ewsxml.FindItemResponseMessage
	}
}

type CreateItemCalendarItemsOperation struct {
	header     ewsxml.Header
	CreateItem struct {
		ewsxml.CreateItem
	}
}

type CreateCalendarItemsResponse struct {
	XMLName          xml.Name `xml:"CreateItemResponse"`
	ResponseMessages struct {
		XMLName                   xml.Name `xml:"ResponseMessages"`
		CreateItemResponseMessage ewsxml.CreateItemResponseMessage
	}
}

func (op *CreateItemCalendarItemsOperation) Header() *ewsxml.Header { return &op.header }
func (op *CreateItemCalendarItemsOperation) Body() interface{}      { return op.CreateItem }

func GetCalendars(ctx context.Context, req Requester, op *FindItemCalendarViewOperation) (*FindItemCalendarViewResponse, error) {
	if op.FindItem.Traversal == "" {
		op.FindItem.Traversal = ewsxml.Traversal_Shallow
	}
	if op.FindItem.ItemShape.BaseShape == "" {
		op.FindItem.ItemShape.BaseShape = ewsxml.BaseShape_Default
	}
	op.FindItem.ParentFolderIds.DistinguishedFolderId.Id = "calendar"

	var out FindItemCalendarViewResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func CreateCalendarItems(ctx context.Context, req Requester, op *CreateItemCalendarItemsOperation) (*CreateCalendarItemsResponse, error) {
	if op.CreateItem.SendMeetingInvitations == "" {
		op.CreateItem.SendMeetingInvitations = string(ewsxml.SendMeetingInvitations_SendToNone)
	}

	var out CreateCalendarItemsResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
