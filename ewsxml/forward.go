package ewsxml

import "encoding/xml"

// The ForwardItem element contains an Exchange store item to forward to recipients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwarditem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/message-ex15websvcsotherref
type SendForwardItem struct {
	XMLName                    xml.Name `xml:"t:ForwardItem"`
	Subject                    string   `xml:"t:Subject"`
	Body                       *SendBody
	ToRecipients               []Recepients  `xml:"t:ToRecipients,omitempty"`
	CcRecipients               []Recepients  `xml:"t:CcRecepients,omitempty"`
	BccRecipients              []Recepients  `xml:"t:BccRecepients,omitempty"`
	IsReadReceiptRequested     bool          `xml:"t:IsReadReceiptRequested"`
	IsDeliveryReceiptRequested bool          `xml:"t:IsDeliveryReceiptRequested"`
	From                       []SendMailbox `xml:"t:From"`
	ReferenceItemId            ReferenceItemId
	NewBodyContent             *NewBodyContent
	//ReceivedBy                 string
	//ReceivedRepresenting       string
}

type NewBodyContent struct {
	XMLName  xml.Name `xml:"t:NewBodyContent"`
	BodyType BodyType `xml:"BodyType,attr"`
	Contents []byte   `xml:",chardata"`
}

type Recepients struct {
	Mailbox SendMailbox
}
