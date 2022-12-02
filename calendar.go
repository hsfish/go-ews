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
		ParentFolderIds ewsxml.SendParentFolderIds
	}
}

func (op *FindItemCalendarViewOperation) Header() *ewsxml.Header { return &op.header }
func (op *FindItemCalendarViewOperation) Body() interface{}      { return op.FindItem }

type FindItemResponse struct {
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

type RespondToCalendarItemOperation struct {
	header     ewsxml.Header
	CreateItem struct {
		ewsxml.CreateItem
	}
}

type RespondToCalendarItemResponse struct {
	XMLName          xml.Name `xml:"CreateItemResponse"`
	ResponseMessages struct {
		XMLName                   xml.Name `xml:"ResponseMessages"`
		CreateItemResponseMessage ewsxml.CreateItemResponseMessage
	}
}

func (op *RespondToCalendarItemOperation) Header() *ewsxml.Header { return &op.header }
func (op *RespondToCalendarItemOperation) Body() interface{}      { return op.CreateItem }

type GetRemindersOperation struct {
	header       ewsxml.Header
	GetReminders struct {
		ewsxml.GetReminders
	}
}

type GetRemindersResponse struct {
	XMLName   xml.Name `xml:"GetRemindersResponse"`
	Reminders ewsxml.Reminders
}

func (op *GetRemindersOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetRemindersOperation) Body() interface{}      { return op.GetReminders }

type SubscribeOperation struct {
	header    ewsxml.Header
	Subscribe ewsxml.Subscribe
}

type SubscribeResponse struct {
	XMLName          xml.Name `xml:"SubscribeResponse"`
	ResponseMessages struct {
		XMLName                  xml.Name `xml:"ResponseMessages"`
		SubscribeResponseMessage ewsxml.SubscribeResponseMessage
	}
}

type SendNotificationResponse struct {
	XMLName          xml.Name `xml:"SendNotification"`
	ResponseMessages struct {
		XMLName                         xml.Name `xml:"ResponseMessages"`
		SendNotificationResponseMessage ewsxml.SendNotificationResponseMessage
	}
}

func (op *SubscribeOperation) Header() *ewsxml.Header { return &op.header }
func (op *SubscribeOperation) Body() interface{}      { return op.Subscribe }

type SendNotificationResultOperation struct {
	header                 ewsxml.Header
	SendNotificationResult ewsxml.SendNotificationResult
}

func (op *SendNotificationResultOperation) Header() *ewsxml.Header { return &op.header }
func (op *SendNotificationResultOperation) Body() interface{}      { return op.SendNotificationResult }

type GetUserConfigurationOperation struct {
	header               ewsxml.Header
	GetUserConfiguration ewsxml.GetUserConfiguration
}
type GetUserConfigurationResponse struct {
	XMLName          xml.Name `xml:"GetUserConfigurationResponse"`
	ResponseMessages struct {
		XMLName                             xml.Name `xml:"ResponseMessages"`
		GetUserConfigurationResponseMessage ewsxml.GetUserConfigurationResponseMessage
	}
}

func (op *GetUserConfigurationOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetUserConfigurationOperation) Body() interface{}      { return op.GetUserConfiguration }

type UpdateUserConfigurationOperation struct {
	header                  ewsxml.Header
	UpdateUserConfiguration ewsxml.UpdateUserConfiguration
}

type UpdateUserConfigurationResponse struct {
	XMLName          xml.Name `xml:"UpdateUserConfigurationResponse"`
	ResponseMessages struct {
		XMLName                                xml.Name `xml:"ResponseMessages"`
		UpdateUserConfigurationResponseMessage ewsxml.Response
	}
}

func (op *UpdateUserConfigurationOperation) Header() *ewsxml.Header { return &op.header }
func (op *UpdateUserConfigurationOperation) Body() interface{}      { return op.UpdateUserConfiguration }

type FindItemQueryStringOperation struct {
	header   ewsxml.Header
	FindItem struct {
		ewsxml.FindItem
		// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/querystring-querystringtype
		QueryString string `xml:"m:QueryString"`
		SortOrder   ewsxml.SortOrder

		ParentFolderIds struct {
			DistinguishedFolderId ewsxml.DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
		} `xml:"m:ParentFolderIds"`
	}
}

func (op *FindItemQueryStringOperation) Header() *ewsxml.Header { return &op.header }
func (op *FindItemQueryStringOperation) Body() interface{}      { return op.FindItem }

type GetFolderOperation struct {
	header    ewsxml.Header
	GetFolder ewsxml.GetFolder
}

type GetFolderResponse struct {
	XMLName          xml.Name `xml:"GetFolderResponse"`
	ResponseMessages struct {
		XMLName                  xml.Name `xml:"ResponseMessages"`
		GetFolderResponseMessage []ewsxml.GetFolderResponseMessage
	}
}

func (op *GetFolderOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetFolderOperation) Body() interface{}      { return op.GetFolder }

type FindFolderOperation struct {
	header     ewsxml.Header
	FindFolder ewsxml.FindFolder
}

type FindFolderResponse struct {
	XMLName          xml.Name `xml:"FindFolderResponse"`
	ResponseMessages struct {
		XMLName                   xml.Name `xml:"ResponseMessages"`
		FindFolderResponseMessage []ewsxml.FindFolderResponseMessage
	}
}

func (op *FindFolderOperation) Header() *ewsxml.Header { return &op.header }
func (op *FindFolderOperation) Body() interface{}      { return op.FindFolder }

type CreateFolderOperation struct {
	header       ewsxml.Header
	CreateFolder ewsxml.CreateFolder
}

type CreateFolderResponse struct {
	XMLName          xml.Name `xml:"CreateFolderResponse"`
	ResponseMessages struct {
		XMLName                     xml.Name `xml:"ResponseMessages"`
		CreateFolderResponseMessage []ewsxml.CreateFolderResponseMessage
	}
}

func (op *CreateFolderOperation) Header() *ewsxml.Header { return &op.header }
func (op *CreateFolderOperation) Body() interface{}      { return op.CreateFolder }

type DeleteFolderOperation struct {
	header       ewsxml.Header
	DeleteFolder ewsxml.DeleteFolder
}

type DeleteFolderResponse struct {
	XMLName          xml.Name `xml:"DeleteFolderResponse"`
	ResponseMessages struct {
		XMLName                     xml.Name `xml:"ResponseMessages"`
		DeleteFolderResponseMessage []ewsxml.DeleteFolderResponseMessage
	}
}

func (op *DeleteFolderOperation) Header() *ewsxml.Header { return &op.header }
func (op *DeleteFolderOperation) Body() interface{}      { return op.DeleteFolder }

func GetCalendars(ctx context.Context, req Requester, op *FindItemCalendarViewOperation) (*FindItemResponse, error) {
	if op.FindItem.Traversal == "" {
		op.FindItem.Traversal = ewsxml.Traversal_Shallow
	}
	if op.FindItem.ItemShape.BaseShape == "" {
		op.FindItem.ItemShape.BaseShape = ewsxml.BaseShape_Default
	}

	if len(op.FindItem.ParentFolderIds.DistinguishedFolderId) == 0 && len(op.FindItem.ParentFolderIds.FolderId) == 0 {
		id := ewsxml.SendDistinguishedFolderId{Id: "calendar"}
		op.FindItem.ParentFolderIds.DistinguishedFolderId = append(op.FindItem.ParentFolderIds.DistinguishedFolderId, id)
	}

	var out FindItemResponse
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

func RespondToCalendarItem(ctx context.Context, req Requester, op *RespondToCalendarItemOperation) (*RespondToCalendarItemResponse, error) {
	if op.CreateItem.MessageDisposition == "" {
		op.CreateItem.MessageDisposition = ewsxml.MessageDisposition_SendAndSaveCopy
	}
	var out RespondToCalendarItemResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func GetCalendarReminders(ctx context.Context, req Requester, op *GetRemindersOperation) (*GetRemindersResponse, error) {
	if op.GetReminders.ReminderType == "" {
		op.GetReminders.ReminderType = "All"
	}
	var out GetRemindersResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func SubscribeToNotifications(ctx context.Context, req Requester, op *SubscribeOperation) (*SubscribeResponse, error) {
	var out SubscribeResponse
	if op.Subscribe.PushSubscriptionRequest.FolderIds.DistinguishedFolderId == nil {
		op.Subscribe.PushSubscriptionRequest.FolderIds.DistinguishedFolderId = []ewsxml.DistinguishedFolderId{{Id: "calendar"}}
	}
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func GetUserConfiguration(ctx context.Context, req Requester, op *GetUserConfigurationOperation) (*GetUserConfigurationResponse, error) {
	var out GetUserConfigurationResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func UpdateUserConfiguration(ctx context.Context, req Requester, op *UpdateUserConfigurationOperation) (*UpdateUserConfigurationResponse, error) {
	var out UpdateUserConfigurationResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func SearchCalendarItem(ctx context.Context, req Requester, op *FindItemQueryStringOperation) (*FindItemResponse, error) {
	if op.FindItem.Traversal == "" {
		op.FindItem.Traversal = ewsxml.Traversal_Shallow
	}
	if op.FindItem.ItemShape.BaseShape == "" {
		op.FindItem.ItemShape.BaseShape = ewsxml.BaseShape_Default
	}
	op.FindItem.ParentFolderIds.DistinguishedFolderId.Id = "calendar"

	var out FindItemResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func GetCalendarFolders(ctx context.Context, req Requester, op *GetFolderOperation) (*GetFolderResponse, error) {
	var out GetFolderResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func FindCalendarFolders(ctx context.Context, req Requester, op *FindFolderOperation) (*FindFolderResponse, error) {
	if op.FindFolder.Traversal == "" {
		op.FindFolder.Traversal = ewsxml.Traversal_Shallow
	}
	var out FindFolderResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func CreateCalendarFolder(ctx context.Context, req Requester, op *CreateFolderOperation) (*CreateFolderResponse, error) {
	if op.CreateFolder.ParentFolderId == nil {
		id := &ewsxml.SendDistinguishedFolderId{
			Id: "calendar",
		}
		op.CreateFolder.ParentFolderId = &ewsxml.SendParentFolderId{DistinguishedFolderId: id}
	}

	var out CreateFolderResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

func DeleteFolder(ctx context.Context, req Requester, op *DeleteFolderOperation) (*DeleteFolderResponse, error) {
	if op.DeleteFolder.DeleteType == "" {
		op.DeleteFolder.DeleteType = ewsxml.DeleteType_HardDelete
	}
	var out DeleteFolderResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
