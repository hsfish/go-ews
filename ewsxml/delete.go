package ewsxml

import (
	"encoding/xml"
)

// DeleteType Describes how an item is deleted. This attribute is required.
type DeleteType string

func (s DeleteType) String() string { return string(s) }

// SendMeetingCancellations Describes whether a calendar item
// deletion is communicated to attendees. This attribute is required when
// calendar items are deleted. This attribute is optional if non-calendar
// items are deleted.
type SendMeetingCancellations string

func (s SendMeetingCancellations) String() string { return string(s) }

// AffectedTaskOccurrences Describes whether a task instance or a task master
// is deleted by a DeleteItem operation. This attribute is required when
// tasks are deleted. This attribute is optional when non-task items are deleted.
type AffectedTaskOccurrences string

func (s AffectedTaskOccurrences) String() string { return string(s) }

// SuppressReadReceipts Indicates whether read receipts for the deleted item
// are suppressed. A text value of true, indicates that the read receipts are
// suppressed. A value of false indicates that the read receipts are sent to
// the sender. This attribute is optional.
type SuppressReadReceipts bool

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// An item is permanently removed from the store.
	DeleteType_HardDelete DeleteType = "HardDelete"
	// An item is moved to the dumpster if the dumpster is enabled.
	DeleteType_SoftDelete DeleteType = "SoftDelete"
	// An item is moved to the Deleted Items folder.
	DeleteType_MoveToDeletedItems DeleteType = "MoveToDeletedItems"

	// A calendar item is deleted without sending a cancellation message.
	SendMeetingCancellations_SendToNone SendMeetingCancellations = "SendToNone"
	// A calendar item is deleted and a cancellation message is sent to all attendees.
	SendMeetingCancellations_SendOnlyToAll SendMeetingCancellations = "SendOnlyToAll"
	// A calendar item is deleted and a cancellation message is sent to all attendees.
	// A copy of the cancellation message is saved in the Sent Items folder.
	SendMeetingCancellations_SendToAllAndSaveCopy SendMeetingCancellations = "SendToAllAndSaveCopy"

	// A delete item request deletes the master task, and therefore all recurring
	// tasks that are associated with the master task.
	AffectedTaskOccurrences_AllOccurrences AffectedTaskOccurrences = "AllOccurrences"
	// A delete item request deletes only specific occurrences of a task.
	AffectedTaskOccurrences_SpecifiedOccurrenceOnly AffectedTaskOccurrences = "SpecifiedOccurrenceOnly"
)

// The DeleteItem element defines a request to delete an item from a mailbox
// in the Exchange store.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitem
type DeleteItem struct {
	XMLName                  xml.Name                 `xml:"m:DeleteItem"`
	DeleteType               DeleteType               `xml:",attr,omitempty"`
	SendMeetingCancellations SendMeetingCancellations `xml:",attr,omitempty"`
	AffectedTaskOccurrences  AffectedTaskOccurrences  `xml:",attr,omitempty"`
	SuppressReadReceipts     SuppressReadReceipts     `xml:",attr"`

	ItemIds *ItemIds
}

// The DeleteItemResponseMessage element contains the status and result
// of a single DeleteItem operation request.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemresponsemessage
type DeleteItemResponseMessage struct {
	XMLName xml.Name `xml:"DeleteItemResponseMessage"`
	Response
	DescriptiveLinkKey int `xml:"DescriptiveLinkKey,attr"`
	RootFolder         RootFolder
}

// The DeleteFolder operation deletes folders from a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolder-operation
type DeleteFolder struct {
	XMLName    xml.Name   `xml:"m:DeleteFolder"`
	DeleteType DeleteType `xml:"DeleteType,attr"`
	FolderIds  SendFolderIds
}

// The DeleteFolderResponseMessage element contains the status and result
// of a single DeleteFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolderresponsemessage
type DeleteFolderResponseMessage struct {
	XMLName xml.Name `xml:"DeleteFolderResponseMessage"`
	Response
	DescriptiveLinkKey int `xml:"DescriptiveLinkKey,attr"`
	RootFolder         RootFolder
}
