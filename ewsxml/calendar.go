package ewsxml

import (
	"strings"
	"time"
)

// The Sensitivity element indicates the sensitivity level of an item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/sensitivity
type Sensitivity string

// The LegacyFreeBusyStatus element represents the free/busy status of the
// calendar item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/legacyfreebusystatus
type LegacyFreeBusyStatus string

// The CalendarItemType element represents the type of a calendar item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritemtype
type CalendarItemType string

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
	// MimeContent                  string      `xml:"t:MimeContent"`
	ItemId         ItemId `xml:"t:ItemId,omitempty"`
	ParentFolderId ItemId `xml:"t:ParentFolderId,omitempty"`
	// ItemClass                    string      `xml:"t:ItemClass"`
	Subject     string      `xml:"t:Subject"`
	Sensitivity Sensitivity `xml:"t:Sensitivity"`
	Body        Body        `xml:"t:Body"`
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
	ReminderIsSet              bool               `xml:"t:ReminderIsSet"`
	ReminderMinutesBeforeStart int                `xml:"t:ReminderMinutesBeforeStart"`
	DisplayCc                  ConcatenatedString `xml:"t:DisplayCc"`
	DisplayTo                  ConcatenatedString `xml:"t:DisplayTo"`
	HasAttachments             bool               `xml:"t:HasAttachments"`
	// ExtendedProperty             string      `xml:"t:ExtendedProperty"`
	// Culture                      string      `xml:"t:Culture"`
	Start time.Time `xml:"t:Start"`
	End   time.Time `xml:"t:End"`
	// OriginalStart                string      `xml:"t:OriginalStart"`
	IsAllDayEvent        bool                 `xml:"t:IsAllDayEvent"`
	LegacyFreeBusyStatus LegacyFreeBusyStatus `xml:"t:LegacyFreeBusyStatus"`
	Location             string               `xml:"t:Location"`
	// When                         string      `xml:"t:When"`
	// IsMeeting                    string      `xml:"t:IsMeeting"`
	// IsCancelled                  string      `xml:"t:IsCancelled"`
	// IsRecurring                  string      `xml:"t:IsRecurring"`
	// MeetingRequestWasSent        string      `xml:"t:MeetingRequestWasSent"`
	// IsResponseRequested          string      `xml:"t:IsResponseRequested"`
	CalendarItemType CalendarItemType `xml:"t:CalendarItemType"`
	// MyResponseType               string      `xml:"t:MyResponseType"`
	Organizer         OneMailbox `xml:"t:Organizer"`
	RequiredAttendees []Attendee `xml:"t:RequiredAttendees"` // []Attendees
	OptionalAttendees []Attendee `xml:"t:OptionalAttendees"` // []Attendees
	Resources         []Attendee `xml:"t:Resources"`         // []Attendees
	// ConflictingMeetingCount      string      `xml:"t:ConflictingMeetingCount"`
	// AdjacentMeetingCount         string      `xml:"t:AdjacentMeetingCount"`
	// ConflictingMeetings          string      `xml:"t:ConflictingMeetings"`
	// AdjacentMeetings             string      `xml:"t:AdjacentMeetings"`
	// Duration                     string      `xml:"t:Duration"`
	// TimeZone                     string      `xml:"t:TimeZone"`
	// AppointmentReplyTime         string      `xml:"t:AppointmentReplyTime"`
	// AppointmentSequenceNumber    string      `xml:"t:AppointmentSequenceNumber"`
	// AppointmentState             string      `xml:"t:AppointmentState"`
	// Recurrence                   string      `xml:"t:Recurrence"`
	// FirstOccurrence              string      `xml:"t:FirstOccurrence"`
	// LastOccurrence               string      `xml:"t:LastOccurrence"`
	// ModifiedOccurrences          string      `xml:"t:ModifiedOccurrences"`
	// DeletedOccurrences           string      `xml:"t:DeletedOccurrences"`
	// MeetingTimeZone              string      `xml:"t:MeetingTimeZone"`
	// StartTimeZone                string      `xml:"t:StartTimeZone"`
	// EndTimeZone                  string      `xml:"t:EndTimeZone"`
	// ConferenceType               string      `xml:"t:ConferenceType"`
	// AllowNewTimeProposal         string      `xml:"t:AllowNewTimeProposal"`
	// IsOnlineMeeting              string      `xml:"t:IsOnlineMeeting"`
	// MeetingWorkspaceUrl          string      `xml:"t:MeetingWorkspaceUrl"`
	// NetShowUrl                   string      `xml:"t:NetShowUrl"`
	// EffectiveRights              string      `xml:"t:EffectiveRights"`
	// LastModifiedName             string      `xml:"t:LastModifiedName"`
	// LastModifiedTime             string      `xml:"t:LastModifiedTime"`
	IsAssociated bool `xml:"t:IsAssociated"`
	// WebClientReadFormQueryString string      `xml:"t:WebClientReadFormQueryString"`
	// WebClientEditFormQueryString string      `xml:"t:WebClientEditFormQueryString"`
	// ConversationId               string      `xml:"t:ConversationId"`
	// UniqueBody                   string      `xml:"t:UniqueBody"`
}

// ConcatenatedString represents the concatenated display string that is used
// for the contents of the element. Each part represents its own value.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaycc
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayto
type ConcatenatedString string

func (c ConcatenatedString) Split(sep string) []string {
	return strings.Split(string(c), sep)
}
