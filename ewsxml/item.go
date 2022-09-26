package ewsxml

import (
	"encoding/xml"
)

type MessageDisposition string

func (s MessageDisposition) String() string { return string(s) }

type SendMeetingInvitations string

func (s SendMeetingInvitations) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// MessageDisposition_SaveOnly indicates the item is updated and saved back
	// to its current folder.
	MessageDisposition_SaveOnly MessageDisposition = "SaveOnly"
	// MessageDisposition_SendOnly indicates the item is updated and sent but no
	// copy is saved.
	MessageDisposition_SendOnly MessageDisposition = "SendOnly"
	// MessageDisposition_SendAndSaveCopy indicates the item is updated and a
	// copy is saved in the folder identified by the SavedItemFolderId element.
	MessageDisposition_SendAndSaveCopy MessageDisposition = "SendAndSaveCopy"

	// SendMeetingInvitations_SendToNone indicates if the item is a meeting
	// request, it is saved as a calendar item but not sent.
	SendMeetingInvitations_SendToNone SendMeetingInvitations = "SendToNone"
	// SendMeetingInvitations_SendOnlyToAll indicates the meeting request is
	// sent to all attendees but is not saved in the Sent Items folder.
	SendMeetingInvitations_SendOnlyToAll SendMeetingInvitations = "SendOnlyToAll"
	// SendMeetingInvitations_SendToAllAndSaveCopy indicates the meeting
	// request is sent to all attendees and a copy is saved in the folder that
	// is identified by the SavedItemFolderId element.
	SendMeetingInvitations_SendToAllAndSaveCopy SendMeetingInvitations = "SendToAllAndSaveCopy"
)

// The SavedItemFolderId element identifies the target folder for operations
// that update, send, and create items in a mailbox.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/saveditemfolderid
type SavedItemFolderId struct {
	XMLName               xml.Name `xml:"m:SavedItemFolderId"`
	FolderId              FolderId
	DistinguishedFolderId DistinguishedFolderId
}

// The CreateItem element defines a request to create an item in the Exchange
// store.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem
type CreateItem struct {
	XMLName                xml.Name           `xml:"m:CreateItem"`
	MessageDisposition     MessageDisposition `xml:"MessageDisposition,attr,omitempty"`
	SendMeetingInvitations string             `xml:"SendMeetingInvitations,attr,omitempty"`
	SavedItemFolderId      *SavedItemFolderId `xml:",omitempty"`
	Items                  MessageItems
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/items
type Items struct {
	// Item                Item                `xml:"t:Item"`
	XMLName      xml.Name  `xml:"Items"`
	Message      []Message `xml:",omitempty"`
	CalendarItem []CalendarItem
	// Contact             Contact             `xml:"t:Contact"`
	// DistributionList    DistributionList    `xml:"t:DistributionList"`
	// MeetingMessage      MeetingMessage      `xml:"t:MeetingMessage"`
	// MeetingRequest      MeetingRequest      `xml:"t:MeetingRequest"`
	// MeetingResponse     MeetingResponse     `xml:"t:MeetingResponse"`
	// MeetingCancellation MeetingCancellation `xml:"t:MeetingCancellation"`
	// Task                Task                `xml:"t:Task"`
	// PostItem            PostItem            `xml:"t:PostItem"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/items
type MessageItems struct {
	XMLName               xml.Name                    `xml:"m:Items"`
	Message               []Message                   `xml:",omitempty"`
	CalendarItem          []SendCalendarItem          `xml:",omitempty"`
	CancelCalendarItem    []SendCancelCalendarItem    `xml:",omitempty"`
	AcceptItem            []SendAcceptItem            `xml:",omitempty"`
	TentativelyAcceptItem []SendTentativelyAcceptItem `xml:",omitempty"`
	DeclineItem           []SendDeclineItem           `xml:",omitempty"`
	// Contact             Contact             `xml:"t:Contact"`
	// DistributionList    DistributionList    `xml:"t:DistributionList"`
	// MeetingMessage      MeetingMessage      `xml:"t:MeetingMessage"`
	// MeetingRequest      MeetingRequest      `xml:"t:MeetingRequest"`
	// MeetingResponse     MeetingResponse     `xml:"t:MeetingResponse"`
	// MeetingCancellation MeetingCancellation `xml:"t:MeetingCancellation"`
	// Task                Task                `xml:"t:Task"`
	// PostItem            PostItem            `xml:"t:PostItem"`
}

type SendItem struct {
	XMLName          xml.Name `xml:"m:SendItem"`
	SaveItemToFolder bool     `xml:"SaveItemToFolder,attr"`

	ItemIds           ItemIds
	SavedItemFolderId SavedItemFolderId
}

type ItemIds struct {
	XMLName               xml.Name `xml:"m:ItemIds"`
	ItemId                []ItemId `xml:"t:ItemId"`
	OccurrenceItemId      []OccurrenceItemId
	RecurringMasterItemId []RecurringMasterItemId
}

// The ItemId element contains the unique identifier and change key of an item
// in the Exchange store.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemid
type ItemId struct {
	Id        string `xml:"Id,attr"`
	ChangeKey string `xml:"ChangeKey,attr,omitempty"`
}

// The OccurrenceItemId element identifies a single occurrence of a recurring
// item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/occurrenceitemid
type OccurrenceItemId struct {
	XMLName           xml.Name `xml:"t:OccurrenceItemId"`
	RecurringMasterId string   `xml:"RecurringMasterId,attr"`
	ChangeKey         string   `xml:"ChangeKey,attr"`
	InstanceIndex     uint     `xml:"InstanceIndex,attr"`
}

// The RecurringMasterItemId element identifies a recurrence master item by
// identifying the identifiers of one of its related occurrence items.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringmasteritemid
type RecurringMasterItemId struct {
	XMLName      xml.Name `xml:"t:RecurringMasterItemId"`
	OccurrenceId string   `xml:"OccurrenceId,attr"`
	ChangeKey    string   `xml:"ChangeKey,attr,omitempty"`
}
