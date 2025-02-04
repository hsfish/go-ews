package ewsxml

import (
	"encoding/xml"
)

type ConflictResolution string

func (s ConflictResolution) String() string { return string(s) }

type SendMeetingInvitationsOrCancellations string

func (s SendMeetingInvitationsOrCancellations) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// ConflictResolution_NeverOverwrite indicates if there is a conflict, the
	// update operation fails and an error is returned.
	ConflictResolution_NeverOverwrite ConflictResolution = "NeverOverwrite"
	// ConflictResolution_AutoResolve indicates the update operation
	// automatically resolves any conflict.
	ConflictResolution_AutoResolve ConflictResolution = "AutoResolve"
	// ConflictResolution_AlwaysOverwrite indicates if there is a conflict, the
	// update operation will overwrite information.
	ConflictResolution_AlwaysOverwrite ConflictResolution = "AlwaysOverwrite"

	// SendMeetingInvitationsOrCancellations_SendToNone indicates the calendar
	// item is updated but updates are not sent to attendees.
	SendMeetingInvitationsOrCancellations_SendToNone SendMeetingInvitationsOrCancellations = "SendToNone"
	// SendMeetingInvitationsOrCancellations_SendOnlyToAll indicates the
	// calendar item is updated and the meeting update is sent to all attendees
	// but is not saved in the Sent Items folder.
	SendMeetingInvitationsOrCancellations_SendOnlyToAll SendMeetingInvitationsOrCancellations = "SendOnlyToAll"
	// SendMeetingInvitationsOrCancellations_SendOnlyToChanged indicates the
	// calendar item is updated and the meeting update is sent only to
	// attendees that are affected by the change in the meeting.
	SendMeetingInvitationsOrCancellations_SendOnlyToChanged SendMeetingInvitationsOrCancellations = "SendOnlyToChanged"
	// SendMeetingInvitationsOrCancellations_SendToAllAndSaveCopy indicates the
	// calendar item is updated, the meeting update is sent to all attendees,
	// and a copy is saved in the Sent Items folder.
	SendMeetingInvitationsOrCancellations_SendToAllAndSaveCopy SendMeetingInvitationsOrCancellations = "SendToAllAndSaveCopy"
	// SendMeetingInvitationsOrCancellations_SendToChangedAndSaveCopy indicates
	// the calendar item is updated, the meeting update is sent to all
	// attendees that are affected by the change in the meeting, and a copy is
	// saved in the Sent Items folder.
	SendMeetingInvitationsOrCancellations_SendToChangedAndSaveCopy SendMeetingInvitationsOrCancellations = "SendToChangedAndSaveCopy"
)

// The UpdateItem element defines a request to update an item in a mailbox.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitem
type UpdateItem struct {
	XMLName                               xml.Name                              `xml:"m:UpdateItem"`
	ConflictResolution                    ConflictResolution                    `xml:"ConflictResolution,attr"`
	MessageDisposition                    MessageDisposition                    `xml:"MessageDisposition,attr"`
	SendMeetingInvitationsOrCancellations SendMeetingInvitationsOrCancellations `xml:"SendMeetingInvitationsOrCancellations,attr"`
	SuppressReadReceipts                  bool                                  `xml:"SuppressReadReceipts,attr"`

	SavedItemFolderId *SavedItemFolderId
	ItemChanges       *ItemChanges
}

// The ItemChanges element contains an array of ItemChange elements that identify items and
// and the updates to apply to the items.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemchanges
type ItemChanges struct {
	XMLName    xml.Name `xml:"m:ItemChanges"`
	ItemChange []ItemChange
}

// The ItemChange element contains an item identifier and the updates to apply
// to the item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemchange
type ItemChange struct {
	XMLName               xml.Name `xml:"t:ItemChange"`
	ItemId                SendItemId
	OccurrenceItemId      *OccurrenceItemId
	RecurringMasterItemId *RecurringMasterItemId
	Updates               Updates
}

type SendItemId struct {
	XMLName   xml.Name `xml:"t:ItemId"`
	Id        string   `xml:"Id,attr"`
	ChangeKey string   `xml:"ChangeKey,attr,omitempty"`
}

type ReferenceItemId struct {
	XMLName   xml.Name `xml:"t:ReferenceItemId"`
	Id        string   `xml:"Id,attr"`
	ChangeKey string   `xml:"ChangeKey,attr,omitempty"`
}

type SendBody struct {
	XMLName  xml.Name `xml:"t:Body"`
	BodyType BodyType `xml:"BodyType,attr"`
	Contents []byte   `xml:",chardata"`
}

type SendMailbox struct {
	XMLName      xml.Name     `xml:"t:Mailbox"`
	Name         string       `xml:"t:Name"`
	EmailAddress string       `xml:"t:EmailAddress"`
	RoutingType  *RoutingType `xml:"t:RoutingType,omitempty"`
	MailboxType  *MailboxType `xml:"t:MailboxType,omitempty"`
	ItemId       *ItemId      `xml:"t:ItemId,omitempty"`
}

type SendAttendee struct {
	XMLName xml.Name `xml:"t:Attendee"`
	Mailbox SendMailbox
}

type SendRequiredAttendees struct {
	XMLName  xml.Name `xml:"t:RequiredAttendees"`
	Attendee []SendAttendee
}

type SendOptionalAttendees struct {
	XMLName  xml.Name `xml:"t:OptionalAttendees"`
	Attendee []SendAttendee
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/updates-item
type Updates struct {
	XMLName         xml.Name `xml:"t:Updates"`
	SetItemField    []SetItemField
	DeleteItemField []DeleteItemField `xml:",omitempty"`
}

// The UpdateItemResponseMessage element contains the status and result
// of a single UpdateItem operation request.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitemresponsemessage
type UpdateItemResponseMessage struct {
	XMLName xml.Name `xml:"UpdateItemResponseMessage"`
	Response
	Items Items
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/setitemfield
type SetItemField struct {
	XMLName      xml.Name `xml:"t:SetItemField"`
	FieldURI     FieldURI
	CalendarItem *SendCalendarItem
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemfield
type DeleteItemField struct {
	XMLName  xml.Name `xml:"t:DeleteItemField"`
	FieldURI FieldURI
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduri
type FieldURI struct {
	XMLName  xml.Name `xml:"t:FieldURI"`
	FieldURI string   `xml:"FieldURI,attr"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedfielduri
type IndexedFieldURI struct {
	XMLName    xml.Name `xml:"t:IndexedFieldURI"`
	FieldURI   string   `xml:"FieldURI,attr"`
	FieldIndex string   `xml:"FieldIndex,attr"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedfielduri
type ExtendedFieldURI struct {
	XMLName                    xml.Name `xml:"t:ExtendedFieldURI"`
	DistinguishedPropertySetId string   `xml:"DistinguishedPropertySetId,attr"`
	PropertySetId              string   `xml:"PropertySetId,attr,omitempty"`
	PropertyTag                string   `xml:"PropertyTag,attr,omitempty"`
	PropertyName               string   `xml:"PropertyName,attr,omitempty"`
	PropertyId                 string   `xml:"PropertyId,attr,omitempty"`
	PropertyType               string   `xml:"PropertyType,attr,omitempty"`
}
