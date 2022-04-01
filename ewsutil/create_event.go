package ewsutil

import (
	"time"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

func CreateHTMLEvent(
	c ews.Client,
	to, optional []string,
	subject, body, location string,
	from time.Time,
	duration time.Duration,
) error {
	return createEvent(c, to, optional, subject, body, location, ewsxml.BodyType_HTML, from, duration)
}

// CreateEvent helper method to send Message
func CreateEvent(
	c ews.Client,
	to, optional []string,
	subject, body, location string,
	from time.Time,
	duration time.Duration,
) error {
	return createEvent(c, to, optional, subject, body, location, ewsxml.BodyType_Text, from, duration)
}

func createEvent(
	c ews.Client,
	to, optional []string,
	subject, body, location string,
	bodyType ewsxml.BodyType,
	from time.Time,
	duration time.Duration,
) error {

	requiredAttendees := make([]ewsxml.Attendee, len(to))
	for i, tt := range to {
		requiredAttendees[i].Mailbox.EmailAddress = tt
	}

	optionalAttendees := make([]ewsxml.Attendee, len(optional))
	for i, tt := range optional {
		optionalAttendees[i].Mailbox.EmailAddress = tt
	}

	room := make([]ewsxml.Attendee, 1)
	room[0].Mailbox.EmailAddress = location

	var m ewsxml.CalendarItem
	m.Subject = subject
	m.Body.BodyType = bodyType
	m.Body.Contents = []byte(body)
	m.ReminderIsSet = true
	m.ReminderMinutesBeforeStart = 15
	m.Start = from
	m.End = from.Add(duration)
	m.LegacyFreeBusyStatus = ewsxml.LegacyFreeBusyStatus_Busy
	m.Location = location
	m.RequiredAttendees = requiredAttendees
	m.OptionalAttendees = optionalAttendees
	m.Resources = room

	return ews.CreateCalendarItem(c, m)
}
