package ewsxml

import (
	"encoding/xml"
	"time"
)

// The Reminders element specifies the reminders returned in the response to a GetReminders request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminders
type Reminders struct {
	XMLName  xml.Name `xml:"Reminders"`
	Reminder []Reminder
}

// The Reminder element specifies a reminder for a task or a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminder
type Reminder struct {
	XMLName               xml.Name `xml:"Reminder"`
	Subject               string
	Location              string
	ReminderTime          time.Time
	StartDate             time.Time
	EndDate               time.Time
	ItemId                *ItemId `xml:"ItemId,omitempty"`
	RecurringMasterItemId string
	ReminderGroup         string
	UID                   string
}

type GetReminders struct {
	XMLName      xml.Name   `xml:"m:GetReminders"`
	BeginTime    *time.Time `xml:"m:BeginTime,omitempty"`
	EndTime      *time.Time `xml:"m:EndTime,omitempty"`
	MaxItems     uint       `xml:"m:MaxItems,omitempty"`
	ReminderType string     `xml:"m:ReminderType,omitempty"`
}
