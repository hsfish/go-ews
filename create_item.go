package ews

import (
	"encoding/xml"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
)

type XMailbox struct {
	Mailbox []ewsxml.Mailbox `xml:"t:Mailbox"`
}

type createItemResponseBodyEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    createItemResponseBody `xml:"Body"`
}
type createItemResponseBody struct {
	CreateItemResponse CreateItemResponse `xml:"CreateItemResponse"`
}

type CreateItemResponse struct {
	ResponseMessages struct {
		CreateItemResponseMessage ewsxml.Response `xml:"CreateItemResponseMessage"`
	} `xml:"ResponseMessages"`
}

// CreateMessageItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-email-message
func CreateMessageItem(c Requester, m ...ewsxml.Message) error {
	var item ewsxml.CreateItem
	item.MessageDisposition = ewsxml.MessageDisposition_SendAndSaveCopy
	item.SavedItemFolderId.DistinguishedFolderId.Id = "sentitems"
	item.Items.Message = append(item.Items.Message, m...)

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	bb, err := c.Request(xmlBytes)
	if err != nil {
		return err
	}

	if err := checkCreateItemResponseForErrors(bb); err != nil {
		return err
	}

	return nil
}

// CreateCalendarItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-calendar-item
func CreateCalendarItem(c Requester, ci ...ewsxml.CalendarItem) error {
	var item ewsxml.CreateItem
	item.SendMeetingInvitations = "SendToAllAndSaveCopy"
	item.SavedItemFolderId.DistinguishedFolderId.Id = "calendar"
	item.Items.CalendarItem = append(item.Items.CalendarItem, ci...)

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	bb, err := c.Request(xmlBytes)
	if err != nil {
		return err
	}

	if err := checkCreateItemResponseForErrors(bb); err != nil {
		return err
	}

	return nil
}

func checkCreateItemResponseForErrors(bb []byte) error {
	var soapResp createItemResponseBodyEnvelop
	if err := xml.Unmarshal(bb, &soapResp); err != nil {
		return err
	}

	resp := soapResp.Body.CreateItemResponse.ResponseMessages.CreateItemResponseMessage
	if resp.ResponseClass == ewsxml.ResponseClass_Error {
		return errors.New(resp.MessageText)
	}
	return nil
}
