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
		FindItemResponseMessage ewsxml.FindItemResponseMessage
	}
}

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
