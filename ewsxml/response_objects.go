package ewsxml

import (
	"encoding/xml"
)

type ResponseObjects struct {
	XMLName xml.Name `xml:"ResponseObjects"`

	AcceptItem              *AcceptItem
	TentativelyAcceptItem   *TentativelyAcceptItem
	DeclineItem             *DeclineItem
	ProposeNewTime          *ProposeNewTime
	ReplyToItem             *ReplyToItem
	ForwardItem             *ForwardItem
	ReplyAllToItem          *ReplyAllToItem
	CancelCalendarItem      *CancelCalendarItem
	RemoveItem              *RemoveItem
	PostReplyItem           *PostReplyItem
	SuppressReadReceipt     *SuppressReadReceipt
	AcceptSharingInvitation *AcceptSharingInvitation
}

type AcceptItem struct {
	XMLName xml.Name `xml:"AcceptItem"`
}

type TentativelyAcceptItem struct {
	XMLName xml.Name `xml:"TentativelyAcceptItem"`
}
type DeclineItem struct {
	XMLName xml.Name `xml:"DeclineItem"`
}
type ProposeNewTime struct {
	XMLName xml.Name `xml:"ProposeNewTime"`
}
type ReplyToItem struct {
	XMLName xml.Name `xml:"ReplyToItem"`
}
type ForwardItem struct {
	XMLName xml.Name `xml:"ForwardItem"`
}
type ReplyAllToItem struct {
	XMLName xml.Name `xml:"ReplyAllToItem"`
}
type CancelCalendarItem struct {
	XMLName xml.Name `xml:"CancelCalendarItem"`
}
type RemoveItem struct {
	XMLName xml.Name `xml:"RemoveItem"`
}
type PostReplyItem struct {
	XMLName xml.Name `xml:"PostReplyItem"`
}
type SuppressReadReceipt struct {
	XMLName xml.Name `xml:"SuppressReadReceipt"`
}
type AcceptSharingInvitation struct {
	XMLName xml.Name `xml:"AcceptSharingInvitation"`
}

type SendAcceptItem struct {
	XMLName         xml.Name `xml:"t:AcceptItem"`
	ReferenceItemId ReferenceItemId
	//ItemClass
	//Sensitivity
	//Body
	//Attachments
	//InternetMessageHeaders
	//Sender
	//ToRecipients
	//CcRecipients
	//BccRecipients
	//IsReadReceiptRequested
	//IsDeliveryReceiptRequested
	//ReceivedBy
	//ReceivedRepresenting
	//ProposedStart
	//ProposedEnd
}

type SendTentativelyAcceptItem struct {
	XMLName         xml.Name `xml:"t:TentativelyAcceptItem"`
	ReferenceItemId ReferenceItemId
}
type SendDeclineItem struct {
	XMLName         xml.Name `xml:"t:DeclineItem"`
	ReferenceItemId ReferenceItemId
}
