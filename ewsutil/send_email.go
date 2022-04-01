package ewsutil

import (
	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

// SendEmail helper method to send Message
func SendEmail(c ews.Client, to []string, subject, body string) error {
	var m ewsxml.Message
	m.ItemClass = "IPM.Note"
	m.Subject = subject
	m.Body.BodyType = ewsxml.BodyType_Text
	m.Body.Contents = []byte(body)
	m.Sender.Mailbox.EmailAddress = c.Username()

	for _, addr := range to {
		m.ToRecipients = append(m.ToRecipients, ewsxml.Mailbox{
			EmailAddress: addr,
		})
	}

	return ews.CreateMessageItem(c, m)
}
