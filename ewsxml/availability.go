package ewsxml

import "encoding/xml"

// The RequestedView element defines the type of calendar information
// that a client requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestedview
type RequestedView string

func (s RequestedView) String() string { return string(s) }

// The FreeBusyViewType element represents the type of free/busy information
// returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyviewtype
type FreeBusyViewType string

func (s FreeBusyViewType) String() string { return string(s) }

// The MinimumSuggestionQuality element defines the quality of meeting
// suggestions to be returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/minimumsuggestionquality
type MinimumSuggestionQuality string

func (s MinimumSuggestionQuality) String() string { return string(s) }

// The AttendeeType element represents the type of attendee that is
// identified in the Email (EmailAddressType) element. This element
// is used in requests for meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendeetype
type AttendeeType string

func (s AttendeeType) String() string { return string(s) }

const (
	RequestedView_None           RequestedView = "None"
	RequestedView_MergedOnly     RequestedView = "MergedOnly"
	RequestedView_FreeBusy       RequestedView = "FreeBusy"
	RequestedView_FreeBusyMerged RequestedView = "FreeBusyMerged"
	RequestedView_Detailed       RequestedView = "Detailed"
	RequestedView_DetailedMerged RequestedView = "DetailedMerged"

	FreeBusyViewType_None           FreeBusyViewType = "None"
	FreeBusyViewType_MergedOnly     FreeBusyViewType = "MergedOnly"
	FreeBusyViewType_FreeBusy       FreeBusyViewType = "FreeBusy"
	FreeBusyViewType_FreeBusyMerged FreeBusyViewType = "FreeBusyMerged"
	FreeBusyViewType_Detailed       FreeBusyViewType = "Detailed"
	FreeBusyViewType_DetailedMerged FreeBusyViewType = "DetailedMerged"

	MinimumSuggestionQuality_Excellent MinimumSuggestionQuality = "Excellent"
	MinimumSuggestionQuality_Good      MinimumSuggestionQuality = "Good"
	MinimumSuggestionQuality_Fair      MinimumSuggestionQuality = "Fair"
	MinimumSuggestionQuality_Poor      MinimumSuggestionQuality = "Poor"

	AttendeeType_Organizer AttendeeType = "Organizer"
	AttendeeType_Required  AttendeeType = "Required"
	AttendeeType_Optional  AttendeeType = "Optional"
	AttendeeType_Room      AttendeeType = "Room"
	AttendeeType_Resource  AttendeeType = "Resource"
)

// The GetUserAvailabilityRequest element contains the arguments used to obtain
// user availability information. This is a root element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailabilityrequest
type GetUserAvailabilityRequest struct {
	XMLName                xml.Name `xml:"m:GetUserAvailabilityRequest"`
	MailboxDataArray       MailboxDataArray
	FreeBusyViewOptions    FreeBusyViewOptions
	SuggestionsViewOptions *SuggestionsViewOptions
}

// The MailboxDataArray element contains a list of mailboxes to query for
// availability information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxdataarray
type MailboxDataArray struct {
	XMLName     xml.Name `xml:"m:MailboxDataArray"`
	MailboxData []MailboxData
}

// The MailboxData element represents an individual mailbox user and options
// for the type of data to be returned about the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxdata
type MailboxData struct {
	XMLName          xml.Name `xml:"t:MailboxData"`
	Email            EmailAddressType
	AttendeeType     AttendeeType `xml:"t:AttendeeType"`
	ExcludeConflicts bool         `xml:"t:ExcludeConflicts"`
}

// The Email element represents the mailbox user for a GetUserAvailability query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/email-emailaddresstype
type EmailAddressType struct {
	XMLName     xml.Name    `xml:"t:Email"`
	Name        string      `xml:"t:Name,omitempty"`
	Address     string      `xml:"t:Address,omitempty"`
	RoutingType RoutingType `xml:"t:RoutingType,omitempty"`
}

// The FreeBusyViewOptions element specifies the type of free/busy information returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyviewoptions
type FreeBusyViewOptions struct {
	XMLName                         xml.Name      `xml:"t:FreeBusyViewOptions"`
	TimeWindow                      TimeWindowStr `xml:"t:TimeWindow"`
	MergedFreeBusyIntervalInMinutes int           `xml:"t:MergedFreeBusyIntervalInMinutes"`
	RequestedView                   RequestedView `xml:"t:RequestedView"`
}

// The SuggestionsViewOptions element contains the options for obtaining meeting
// suggestion information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionsviewoptions
type SuggestionsViewOptions struct {
	XMLName                        xml.Name                 `xml:"t:SuggestionsViewOptions"`
	GoodThreshold                  int                      `xml:"t:GoodThreshold"`
	MaximumResultsByDay            int                      `xml:"t:MaximumResultsByDay"`
	MaximumNonWorkHourResultsByDay int                      `xml:"t:MaximumNonWorkHourResultsByDay"`
	MeetingDurationInMinutes       int                      `xml:"t:MeetingDurationInMinutes"`
	MinimumSuggestionQuality       MinimumSuggestionQuality `xml:"t:MinimumSuggestionQuality"`
	DetailedSuggestionsWindow      TimeWindowStr            `xml:"t:DetailedSuggestionsWindow"`
	//CurrentMeetingTime             string                   `xml:"t:CurrentMeetingTime"`
}

// The FreeBusyResponseArray element contains the requested users' availability
// information and the response status.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyresponsearray
type FreeBusyResponseArray struct {
	XMLName          xml.Name `xml:"FreeBusyResponseArray"`
	FreeBusyResponse []FreeBusyResponse
}

// The FreeBusyResponse element contains the free/busy information for a single
// mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyresponse
type FreeBusyResponse struct {
	Response
	FreeBusyView FreeBusyView
}

// The FreeBusyView element contains availability information for a specific user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyview
type FreeBusyView struct {
	XMLName            xml.Name `xml:"FreeBusyView"`
	FreeBusyViewType   FreeBusyViewType
	MergedFreeBusy     string
	CalendarEventArray struct {
		CalendarEvent []CalendarEvent
	}
	WorkingHours WorkingHours
}

type CalendarEvent struct {
	XMLName              xml.Name `xml:"CalendarEvent"`
	StartTime            string
	EndTime              string
	BusyType             LegacyFreeBusyStatus
	CalendarEventDetails struct {
		ID            string
		Subject       string
		Location      string
		IsMeeting     bool
		IsRecurring   bool
		IsException   bool
		IsReminderSet bool
		IsPrivate     bool
	}
}

// The WorkingHours element represents the time zone settings and working hours
// for the requested mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workinghours-ex15websvcsotherref
type WorkingHours struct {
	XMLName            xml.Name `xml:"WorkingHours"`
	TimeZone           TimeZone
	WorkingPeriodArray struct {
		WorkingPeriod []WorkingPeriod
	}
}

// The WorkingPeriod element contains the work week days and hours of the
// mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workingperiod
type WorkingPeriod struct {
	DayOfWeek          string
	StartTimeInMinutes int
	EndTimeInMinutes   int
}

// The SuggestionsResponse element contains response status information and
// suggestion data for requested meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionsresponse
type SuggestionsResponse struct {
	XMLName                  xml.Name `xml:"SuggestionsResponse"`
	SuggestionDayResultArray struct {
		SuggestionDayResult []SuggestionDayResult
	}
}

type SuggestionDayResultArray struct {
	XMLName xml.Name `xml:"SuggestionDayResultArray"`
}

// The SuggestionDayResult element represents a single day that contains
// suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestiondayresult
type SuggestionDayResult struct {
	XMLName         xml.Name `xml:"SuggestionDayResult"`
	Date            string
	DayQuality      MinimumSuggestionQuality
	SuggestionArray struct {
		Suggestion []Suggestion
	}
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestion
// The Suggestion element represents a single meeting suggestion.
type Suggestion struct {
	XMLName           xml.Name `xml:"Suggestion"`
	MeetingTime       string
	IsWorkTime        bool
	SuggestionQuality MinimumSuggestionQuality
	//AttendeeConflictDataArray
}
