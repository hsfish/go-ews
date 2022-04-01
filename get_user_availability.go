package ews

import (
	"encoding/xml"
	"errors"
	"time"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type AttendeeType string
type BusyType string

//goland:noinspection GoUnusedConst
const (
	AttendeeTypeOrganizer AttendeeType = "Organizer"
	AttendeeTypeRequired  AttendeeType = "Required"
	AttendeeTypeOptional  AttendeeType = "Optional"
	AttendeeTypeRoom      AttendeeType = "Room"
	AttendeeTypeResource  AttendeeType = "Resource"

	RequestedViewNone           = "None"
	RequestedViewMergedOnly     = "MergedOnly"
	RequestedViewFreeBusy       = "FreeBusy"
	RequestedViewFreeBusyMerged = "FreeBusyMerged"
	RequestedViewDetailed       = "Detailed"
	RequestedViewDetailedMerged = "DetailedMerged"

	BusyTypeFree      BusyType = "Free"
	BusyTypeTentative BusyType = "Tentative"
	BusyTypeBusy      BusyType = "Busy"
	BusyTypeOOF       BusyType = "OOF"
	BusyTypeNoData    BusyType = "NoData"
)

type GetUserAvailabilityRequest struct {
	XMLName             struct{}            `xml:"m:GetUserAvailabilityRequest"`
	TimeZone            ewsxml.TimeZone     `xml:"t:TimeZone"`
	MailboxDataArray    MailboxDataArray    `xml:"m:MailboxDataArray"`
	FreeBusyViewOptions FreeBusyViewOptions `xml:"t:FreeBusyViewOptions"`
}

type FreeBusyViewOptions struct {
	TimeWindow                      ewsxml.TimeWindow `xml:"t:TimeWindow"`
	MergedFreeBusyIntervalInMinutes int               `xml:"t:MergedFreeBusyIntervalInMinutes,omitempty"`
	RequestedView                   string            `xml:"t:RequestedView"`
}

type MailboxDataArray struct {
	MailboxData []MailboxData `xml:"t:MailboxData"`
}

type MailboxData struct {
	Email            Email        `xml:"t:Email"`
	AttendeeType     AttendeeType `xml:"t:AttendeeType"`
	ExcludeConflicts bool         `xml:"t:ExcludeConflicts"`
}

type Email struct {
	Name        string `xml:"t:Name"`
	Address     string `xml:"t:Address"`
	RoutingType string `xml:"t:RoutingType"`
}

type GetUserAvailabilityResponse struct {
	FreeBusyResponseArray FreeBusyResponseArray `xml:"FreeBusyResponseArray"`
	SuggestionsResponse   SuggestionsResponse   `xml:"SuggestionsResponse"`
}
type SuggestionsResponse struct {
	ResponseMessage          ResponseMessage          `xml:"ResponseMessage"`
	SuggestionDayResultArray SuggestionDayResultArray `xml:"SuggestionDayResultArray"`
}

type SuggestionDayResultArray struct {
	SuggestionDayResult []SuggestionDayResult `xml:"SuggestionDayResult"`
}
type SuggestionDayResult struct {
	Date            time.Time       `xml:"Date"`
	DayQuality      string          `xml:"DayQuality"`
	SuggestionArray SuggestionArray `xml:"SuggestionArray"`
}

type SuggestionArray struct {
	Suggestion []Suggestion `xml:"Suggestion"`
}

type Suggestion struct {
	MeetingTime                 time.Time                   `xml:"MeetingTime"`
	IsWorkTime                  bool                        `xml:"IsWorkTime"`
	SuggestionQuality           string                      `xml:"SuggestionQuality"`
	ArrayOfAttendeeConflictData ArrayOfAttendeeConflictData `xml:"ArrayOfAttendeeConflictData"`
}

type ArrayOfAttendeeConflictData struct {
	UnknownAttendeeConflictData     string                    `xml:"UnknownAttendeeConflictData"`
	IndividualAttendeeConflictData  string                    `xml:"IndividualAttendeeConflictData"`
	TooBigGroupAttendeeConflictData string                    `xml:"TooBigGroupAttendeeConflictData"`
	GroupAttendeeConflictData       GroupAttendeeConflictData `xml:"GroupAttendeeConflictData"`
}

type GroupAttendeeConflictData struct {
	NumberOfMembers             int `xml:"NumberOfMembers"`
	NumberOfMembersAvailable    int `xml:"NumberOfMembersAvailable"`
	NumberOfMembersWithConflict int `xml:"NumberOfMembersWithConflict"`
	NumberOfMembersWithNoData   int `xml:"NumberOfMembersWithNoData"`
}

type FreeBusyResponseArray struct {
	FreeBusyResponse []FreeBusyResponse `xml:"FreeBusyResponse"`
}

type FreeBusyResponse struct {
	ResponseMessage ResponseMessage `xml:"ResponseMessage"`
	FreeBusyView    FreeBusyView    `xml:"FreeBusyView"`
}

type ResponseMessage struct {
	ewsxml.Response
	DescriptiveLinkKey int `xml:"DescriptiveLinkKey"`
}

type FreeBusyView struct {
	FreeBusyViewType   string             `xml:"FreeBusyViewType"`
	MergedFreeBusy     string             `xml:"MergedFreeBusy"`
	CalendarEventArray CalendarEventArray `xml:"CalendarEventArray"`
	WorkingHours       WorkingHours       `xml:"WorkingHours"`
}

type WorkingHours struct {
	TimeZone           ewsxml.TimeZone    `xml:"TimeZone"`
	WorkingPeriodArray WorkingPeriodArray `xml:"WorkingPeriodArray"`
}

type WorkingPeriodArray struct {
	WorkingPeriod []WorkingPeriod `xml:"WorkingPeriod"`
}

type WorkingPeriod struct {
	DayOfWeek          string `xml:"DayOfWeek"`
	StartTimeInMinutes int    `xml:"StartTimeInMinutes"`
	EndTimeInMinutes   int    `xml:"EndTimeInMinutes"`
}

type CalendarEventArray struct {
	CalendarEvent []CalendarEvent `xml:"CalendarEvent"`
}

type CalendarEvent struct {
	StartTime            ewsxml.Time          `xml:"StartTime"`
	EndTime              ewsxml.Time          `xml:"EndTime"`
	BusyType             BusyType             `xml:"BusyType"`
	CalendarEventDetails CalendarEventDetails `xml:"CalendarEventDetails"`
}

type CalendarEventDetails struct {
	ID            string `xml:"ID"`
	Subject       string `xml:"Subject"`
	Location      string `xml:"Location"`
	IsMeeting     bool   `xml:"IsMeeting"`
	IsRecurring   bool   `xml:"IsRecurring"`
	IsException   bool   `xml:"IsException"`
	IsReminderSet bool   `xml:"IsReminderSet"`
	IsPrivate     bool   `xml:"IsPrivate"`
}

type getUserAvailabilityResponseEnvelop struct {
	XMLName struct{}                        `xml:"Envelope"`
	Body    getUserAvailabilityResponseBody `xml:"Body"`
}
type getUserAvailabilityResponseBody struct {
	GetUserAvailabilityResponse GetUserAvailabilityResponse `xml:"GetUserAvailabilityResponse"`
}

// GetUserAvailability
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailability-operation
func GetUserAvailability(c Requester, r *GetUserAvailabilityRequest) (*GetUserAvailabilityResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.Request(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getUserAvailabilityResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	resp := soapResp.Body.GetUserAvailabilityResponse

	err = checkForFunctionalError(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func checkForFunctionalError(resp *GetUserAvailabilityResponse) error {

	if len(resp.FreeBusyResponseArray.FreeBusyResponse) > 0 {
		for _, rr := range resp.FreeBusyResponseArray.FreeBusyResponse {
			if rr.ResponseMessage.ResponseClass == ewsxml.ResponseClass_Error {
				return errors.New(rr.ResponseMessage.MessageText)
			}
		}
	}
	return nil
}
