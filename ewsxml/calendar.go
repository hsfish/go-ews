package ewsxml

import (
	"encoding/xml"
	"strings"
	"time"
)

// The Sensitivity element indicates the sensitivity level of an item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/sensitivity
type Sensitivity string

func (s Sensitivity) String() string { return string(s) }

// The LegacyFreeBusyStatus element represents the free/busy status of the
// calendar item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/legacyfreebusystatus
type LegacyFreeBusyStatus string

func (s LegacyFreeBusyStatus) String() string { return string(s) }

// The CalendarItemType element represents the type of a calendar item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritemtype
type CalendarItemType string

func (s Sensitivity) CalendarItemType() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	Sensitivity_Normal       Sensitivity = "Normal"
	Sensitivity_Personal     Sensitivity = "Personal"
	Sensitivity_Private      Sensitivity = "Private"
	Sensitivity_Confidential Sensitivity = "Confidential"

	LegacyFreeBusyStatus_Free             LegacyFreeBusyStatus = "Free"
	LegacyFreeBusyStatus_Tentative        LegacyFreeBusyStatus = "Tentative"
	LegacyFreeBusyStatus_Busy             LegacyFreeBusyStatus = "Busy"
	LegacyFreeBusyStatus_OOF              LegacyFreeBusyStatus = "OOF"
	LegacyFreeBusyStatus_WorkingElsewhere LegacyFreeBusyStatus = "WorkingElsewhere"
	LegacyFreeBusyStatus_NoData           LegacyFreeBusyStatus = "NoData"

	// CalendarItemType_Single indicates the item is not associated with a
	// recurring calendar item.
	CalendarItemType_Single CalendarItemType = "Single"
	// CalendarItemType_Occurrence indicates the item is an occurrence of a
	// recurring calendar item.
	CalendarItemType_Occurrence CalendarItemType = "Occurrence"
	// CalendarItemType_Exception indicates the item is an exception to a
	// recurring calendar item.
	CalendarItemType_Exception CalendarItemType = "Exception"
	// CalendarItemType_RecurringMaster indicates the item is master for a set
	// of recurring calendar items.
	CalendarItemType_RecurringMaster CalendarItemType = "RecurringMaster"
)

// The CalendarItem element represents an Exchange calendar item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritem
type CalendarItem struct {
	XMLName xml.Name `xml:"CalendarItem"`
	Data    []byte   `xml:",innerxml"`
	// MimeContent                  string      `xml:"t:MimeContent"`
	ItemId         *ItemId `xml:"ItemId,omitempty"`
	ParentFolderId *ItemId `xml:",omitempty"`
	// ItemClass                    string      `xml:"t:ItemClass"`
	Subject     string      `xml:",omitempty"`
	Sensitivity Sensitivity `xml:",omitempty"`
	Body        *Body       `xml:",omitempty"`
	// Attachments                  string      `xml:"t:Attachments"`
	// DateTimeReceived             string      `xml:"t:DateTimeReceived"`
	// Size                         string      `xml:"t:Size"`
	// Categories                   string      `xml:"t:Categories"`
	// InReplyTo                    string      `xml:"t:InReplyTo"`
	// IsSubmitted                  string      `xml:"t:IsSubmitted"`
	// IsDraft                      string      `xml:"t:IsDraft"`
	// IsFromMe                     string      `xml:"t:IsFromMe"`
	// IsResend                     string      `xml:"t:IsResend"`
	// IsUnmodified                 string      `xml:"t:IsUnmodified"`
	// InternetMessageHeaders       string      `xml:"t:InternetMessageHeaders"`
	// DateTimeSent                 string      `xml:"t:DateTimeSent"`
	// DateTimeCreated              string      `xml:"t:DateTimeCreated"`
	// ResponseObjects              string      `xml:"t:ResponseObjects"`
	// ReminderDueBy                string      `xml:"t:ReminderDueBy"`
	ReminderIsSet              bool
	ReminderMinutesBeforeStart int
	DisplayCc                  ConcatenatedString `xml:",omitempty"`
	DisplayTo                  ConcatenatedString `xml:",omitempty"`
	HasAttachments             bool               `xml:",omitempty"`
	// ExtendedProperty             string      `xml:"t:ExtendedProperty"`
	// Culture                      string      `xml:"t:Culture"`
	Start time.Time
	End   time.Time
	// OriginalStart                string      `xml:"t:OriginalStart"`
	IsAllDayEvent        bool
	LegacyFreeBusyStatus LegacyFreeBusyStatus `xml:",omitempty"`
	Location             string
	// When                         string      `xml:"t:When"`
	IsMeeting   bool
	IsCancelled bool
	IsRecurring bool
	// MeetingRequestWasSent        string      `xml:"t:MeetingRequestWasSent"`
	// IsResponseRequested          string      `xml:"t:IsResponseRequested"`
	CalendarItemType CalendarItemType
	// MyResponseType               string      `xml:"t:MyResponseType"`
	Organizer         *OneMailbox `xml:",omitempty"`
	RequiredAttendees *RequiredAttendees
	OptionalAttendees *OptionalAttendees
	Resources         []Attendee // []Attendees
	// ConflictingMeetingCount      string      `xml:"t:ConflictingMeetingCount"`
	// AdjacentMeetingCount         string      `xml:"t:AdjacentMeetingCount"`
	// ConflictingMeetings          string      `xml:"t:ConflictingMeetings"`
	// AdjacentMeetings             string      `xml:"t:AdjacentMeetings"`
	// Duration                     string      `xml:"t:Duration"`
	// TimeZone                     string      `xml:"t:TimeZone"`
	// AppointmentReplyTime         string      `xml:"t:AppointmentReplyTime"`
	// AppointmentSequenceNumber    string      `xml:"t:AppointmentSequenceNumber"`
	// AppointmentState             string      `xml:"t:AppointmentState"`
	Recurrence *Recurrence
	// FirstOccurrence              string      `xml:"t:FirstOccurrence"`
	// LastOccurrence               string      `xml:"t:LastOccurrence"`
	// ModifiedOccurrences          string      `xml:"t:ModifiedOccurrences"`
	// DeletedOccurrences           string      `xml:"t:DeletedOccurrences"`
	// MeetingTimeZone              string      `xml:"t:MeetingTimeZone"`
	// StartTimeZone                string      `xml:"t:StartTimeZone"`
	// EndTimeZone                  string      `xml:"t:EndTimeZone"`
	StartTimeZoneId string `xml:",omitempty"`
	EndTimeZoneId   string `xml:",omitempty"`
	// ConferenceType               string      `xml:"t:ConferenceType"`
	// AllowNewTimeProposal         string      `xml:"t:AllowNewTimeProposal"`
	// IsOnlineMeeting              string      `xml:"t:IsOnlineMeeting"`
	// MeetingWorkspaceUrl          string      `xml:"t:MeetingWorkspaceUrl"`
	// NetShowUrl                   string      `xml:"t:NetShowUrl"`
	// EffectiveRights              string      `xml:"t:EffectiveRights"`
	// LastModifiedName             string      `xml:"t:LastModifiedName"`
	// LastModifiedTime             string      `xml:"t:LastModifiedTime"`
	IsAssociated bool `xml:",omitempty"`
	// WebClientReadFormQueryString string      `xml:"t:WebClientReadFormQueryString"`
	// WebClientEditFormQueryString string      `xml:"t:WebClientEditFormQueryString"`
	// ConversationId               string      `xml:"t:ConversationId"`
	// UniqueBody                   string      `xml:"t:UniqueBody"`
}
type SendCalendarItem struct {
	XMLName           xml.Name   `xml:"t:CalendarItem"`
	Subject           *string    `xml:"t:Subject,omitempty"`
	Start             *time.Time `xml:"t:Start,omitempty"`
	End               *time.Time `xml:"t:End,omitempty"`
	StartTimeZone     *StartTimeZone
	EndTimeZone       *EndTimeZone
	IsAllDayEvent     *bool   `xml:"t:IsAllDayEvent,omitempty"`
	Location          *string `xml:"t:Location,omitempty"`
	Body              *SendBody
	RequiredAttendees *SendRequiredAttendees
	OptionalAttendees *SendOptionalAttendees
	Recurrence        *SendRecurrence
}

// ConcatenatedString represents the concatenated display string that is used
// for the contents of the element. Each part represents its own value.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaycc
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayto
type ConcatenatedString string

func (c ConcatenatedString) Split(sep string) []string {
	return strings.Split(string(c), sep)
}

func (s ConcatenatedString) String() string { return string(s) }
