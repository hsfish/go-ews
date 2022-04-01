package ews

import (
	"encoding/xml"
	"log"
	"testing"
	"time"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/stretchr/testify/assert"
)

func Test_marshal_CalendarItem(t *testing.T) {
	attendees := make([]ewsxml.Attendee, 0)
	attendees = append(attendees,
		ewsxml.Attendee{Mailbox: ewsxml.Mailbox{EmailAddress: "User1@example.com"}},
		ewsxml.Attendee{Mailbox: ewsxml.Mailbox{EmailAddress: "User2@example.com"}},
	)

	var item ewsxml.CalendarItem
	item.Subject = "Planning Meeting"
	item.Body.BodyType = ewsxml.BodyType_Text
	item.Body.Contents = []byte("Plan the agenda for next week's meeting.")
	item.ReminderIsSet = true
	item.ReminderMinutesBeforeStart = 60
	item.Start, _ = time.Parse(time.RFC3339, "2006-11-02T14:00:00Z")
	item.End, _ = time.Parse(time.RFC3339, "2006-11-02T15:00:00Z")
	item.LegacyFreeBusyStatus = ewsxml.LegacyFreeBusyStatus_Busy
	item.Location = "Conference Room 721"
	item.RequiredAttendees = attendees

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<CalendarItem>
  <t:Subject>Planning Meeting</t:Subject>
  <t:Body BodyType="Text">Plan the agenda for next week&#39;s meeting.</t:Body>
  <t:ReminderIsSet>true</t:ReminderIsSet>
  <t:ReminderMinutesBeforeStart>60</t:ReminderMinutesBeforeStart>
  <t:Start>2006-11-02T14:00:00Z</t:Start>
  <t:End>2006-11-02T15:00:00Z</t:End>
  <t:IsAllDayEvent>false</t:IsAllDayEvent>
  <t:LegacyFreeBusyStatus>Busy</t:LegacyFreeBusyStatus>
  <t:Location>Conference Room 721</t:Location>
  <t:RequiredAttendees>
    <t:Attendee>
      <t:Mailbox>
        <t:EmailAddress>User1@example.com</t:EmailAddress>
      </t:Mailbox>
    </t:Attendee>
    <t:Attendee>
      <t:Mailbox>
        <t:EmailAddress>User2@example.com</t:EmailAddress>
      </t:Mailbox>
    </t:Attendee>
  </t:RequiredAttendees>
</CalendarItem>`, string(xmlBytes))
}
