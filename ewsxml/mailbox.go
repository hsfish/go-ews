package ewsxml

import (
	"encoding/xml"
)

// The RoutingType element represents the routing protocol for the recipient.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/routingtype-emailaddress
type RoutingType string

func (s RoutingType) String() string { return string(s) }

// The MailboxType element represents the type of mailbox that is represented by
// the e-mail address.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxtype
type MailboxType string

func (s MailboxType) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	RoutingType_Smtp RoutingType = "SMTP"
	RoutingType_EX   RoutingType = "EX"

	// MailboxType_Mailbox represents a mail-enabled Active Directory object.
	MailboxType_Mailbox MailboxType = "Mailbox"
	// MailboxType_PublicDL represents a public distribution list.
	MailboxType_PublicDL MailboxType = "PublicDL"
	// MailboxType_PrivateDL represents a private distribution list in a user's
	// mailbox.
	MailboxType_PrivateDL MailboxType = "PrivateDL"
	// MailboxType_Contact represents a contact in a user's mailbox.
	MailboxType_Contact MailboxType = "Contact"
	// MailboxType_PublicFolder represents a public folder.
	MailboxType_PublicFolder MailboxType = "PublicFolder"
	// MailboxType_Unknown represents an unknown type of mailbox.
	MailboxType_Unknown MailboxType = "Unknown"
	// MailboxType_OneOff represents a one-off member of a personal distribution
	// list.
	MailboxType_OneOff MailboxType = "OneOff"
	// MailboxType_GroupMailbox represents a group mailbox.
	MailboxType_GroupMailbox MailboxType = "GroupMailbox"
)

// The Mailbox element identifies a mail-enabled Active Directory object.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox
type Mailbox struct {
	Name         string      `xml:"Name"`
	EmailAddress string      `xml:"EmailAddress"`
	RoutingType  RoutingType `xml:"RoutingType,omitempty"`
	MailboxType  MailboxType `xml:"MailboxType,omitempty"`
	ItemId       ItemId      `xml:"ItemId,omitempty"`
}

// OneMailbox is a wrapper with only a single Mailbox element inside.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/sender
type OneMailbox struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type Attendee struct {
	XMLName xml.Name `xml:"Attendee"`
	Mailbox Mailbox
}

type RequiredAttendees struct {
	XMLName  xml.Name `xml:"RequiredAttendees"`
	Attendee []Attendee
}

type OptionalAttendees struct {
	XMLName  xml.Name `xml:"OptionalAttendees"`
	Attendee []Attendee
}
