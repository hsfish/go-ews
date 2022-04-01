package ews

import (
	"encoding/xml"
	"time"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type FindItemCalendarViewRequest struct {
	ewsxml.FindItem
	CalendarView ewsxml.CalendarView

	// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderids
	ParentFolderIds struct {
		DistinguishedFolderId ewsxml.DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
	} `xml:"m:ParentFolderIds"`
}

type FindItemCalendarViewResponse struct {
	XMLName          xml.Name `xml:"FindItemResponse"`
	ResponseMessages struct {
		FindItemResponseMessage struct {
			ewsxml.FindItemResponseMessage
			ewsxml.RootFolder
		}
	} `xml:"m:ResponseMessages"`
}

func GetCalendars(req Requester, email string, start, end time.Time) (*FindItemCalendarViewResponse, error) {
	var in FindItemCalendarViewRequest
	in.Traversal = ewsxml.Traversal_Shallow
	in.ItemShape.BaseShape = ewsxml.BaseShape_Default
	in.ParentFolderIds.DistinguishedFolderId.Id = "calendar"
	in.ParentFolderIds.DistinguishedFolderId.Mailbox.EmailAddress = email
	in.CalendarView.StartDate = start
	in.CalendarView.EndDate = end
	in.CalendarView.MaxEntriesReturned = 10

	var out FindItemCalendarViewResponse
	err := requestAndUnmarshal(req, in, &out)
	return &out, err
}
