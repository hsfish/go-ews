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

type GetItemCalendarItemsOperation struct {
	header  ewsxml.Header
	GetItem ewsxml.GetItem
}

type GetItemCalendarItemsResponse struct {
	XMLName          xml.Name `xml:"GetItemResponse"`
	ResponseMessages struct {
		XMLName                xml.Name `xml:"ResponseMessages"`
		GetItemResponseMessage ewsxml.GetItemResponseMessage
	}
}

func (op *GetItemCalendarItemsOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetItemCalendarItemsOperation) Body() interface{}      { return op.GetItem }

type UpdateCalendarItemsOperation struct {
	header     ewsxml.Header
	UpdateItem struct {
		ewsxml.UpdateItem
	}
}

type UpdateCalendarItemResponse struct {
	XMLName          xml.Name `xml:"UpdateItemResponse"`
	ResponseMessages struct {
		XMLName                   xml.Name `xml:"ResponseMessages"`
		UpdateItemResponseMessage ewsxml.UpdateItemResponseMessage
	}
}

func (op *UpdateCalendarItemsOperation) Header() *ewsxml.Header { return &op.header }
func (op *UpdateCalendarItemsOperation) Body() interface{}      { return op.UpdateItem }

type DeleteItemCalendarItemsOperation struct {
	header     ewsxml.Header
	DeleteItem ewsxml.DeleteItem
}

type DeleteItemCalendarItemsResponse struct {
	XMLName          xml.Name `xml:"DeleteItemResponse"`
	ResponseMessages struct {
		XMLName                   xml.Name `xml:"ResponseMessages"`
		DeleteItemResponseMessage ewsxml.DeleteItemResponseMessage
	}
}

func (op *DeleteItemCalendarItemsOperation) Header() *ewsxml.Header { return &op.header }
func (op *DeleteItemCalendarItemsOperation) Body() interface{}      { return op.DeleteItem }

type CancelCalendarItemOperation struct {
	header     ewsxml.Header
	CreateItem ewsxml.CreateItem
}

type CancelCalendarItemResponse struct {
	XMLName          xml.Name `xml:"CreateItemResponse"`
	ResponseMessages struct {
		XMLName            xml.Name `xml:"ResponseMessages"`
		CreateItemResponse ewsxml.CreateItemResponseMessage
	}
}

func (op *CancelCalendarItemOperation) Header() *ewsxml.Header { return &op.header }
func (op *CancelCalendarItemOperation) Body() interface{}      { return op.CreateItem }

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

func GetCalendarItems(ctx context.Context, req Requester, op *GetItemCalendarItemsOperation) (*GetItemCalendarItemsResponse, error) {
	if op.GetItem.ItemShape.BaseShape == "" {
		op.GetItem.ItemShape.BaseShape = ewsxml.BaseShape_AllProperties
	}
	var out GetItemCalendarItemsResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func UpdateCalendarItems(ctx context.Context, req Requester, op *UpdateCalendarItemsOperation) (*UpdateCalendarItemResponse, error) {
	if op.UpdateItem.SendMeetingInvitationsOrCancellations == "" {
		op.UpdateItem.SendMeetingInvitationsOrCancellations = ewsxml.SendMeetingInvitationsOrCancellations_SendToNone
	}

	if op.UpdateItem.ConflictResolution == "" {
		op.UpdateItem.ConflictResolution = ewsxml.ConflictResolution_AutoResolve
	}

	if op.UpdateItem.MessageDisposition == "" {
		op.UpdateItem.MessageDisposition = ewsxml.MessageDisposition_SaveOnly
	}
	var out UpdateCalendarItemResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func DeleteCalendarItems(ctx context.Context, req Requester, op *DeleteItemCalendarItemsOperation) (*DeleteItemCalendarItemsResponse, error) {
	if op.DeleteItem.DeleteType == "" {
		op.DeleteItem.DeleteType = ewsxml.DeleteType_MoveToDeletedItems
	}

	if op.DeleteItem.SendMeetingCancellations == "" {
		op.DeleteItem.SendMeetingCancellations = ewsxml.SendMeetingCancellations_SendToAllAndSaveCopy
	}
	var out DeleteItemCalendarItemsResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func CancelCalendarItem(ctx context.Context, req Requester, op *CancelCalendarItemOperation) (*CancelCalendarItemResponse, error) {
	if op.CreateItem.MessageDisposition == "" {
		op.CreateItem.MessageDisposition = ewsxml.MessageDisposition_SendAndSaveCopy
	}
	var out CancelCalendarItemResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
