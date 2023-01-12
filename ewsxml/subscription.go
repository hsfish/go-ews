package ewsxml

import (
	"encoding/xml"
	"time"
)

// The EventType element is used to create a subscription and identifies
// an event type to be reported in a notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventtype
type EventType string

func (s EventType) String() string { return string(s) }

// The SubscriptionStatus element describes the status of a push subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriptionstatus
type SubscriptionStatus string

func (s SubscriptionStatus) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	EventType_CopiedEvent          EventType = "CopiedEvent"
	EventType_CreatedEvent         EventType = "CreatedEvent"
	EventType_DeletedEvent         EventType = "DeletedEvent"
	EventType_ModifiedEvent        EventType = "ModifiedEvent"
	EventType_MovedEvent           EventType = "MovedEvent"
	EventType_NewMailEvent         EventType = "NewMailEvent"
	EventType_FreeBusyChangedEvent EventType = "FreeBusyChangedEvent"

	SubscriptionStatus_OK          SubscriptionStatus = "OK"
	SubscriptionStatus_Unsubscribe SubscriptionStatus = "Unsubscribe"
)

// The Subscribe element contains the properties used to create subscriptions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscribe
type Subscribe struct {
	XMLName xml.Name `xml:"m:Subscribe"`
	//PullSubscriptionRequest PullSubscriptionRequest
	PushSubscriptionRequest *PushSubscriptionRequest
	//StreamingSubscriptionRequest StreamingSubscriptionRequest
}

// The PushSubscriptionRequest element represents a subscription to a push-based
// event notification subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pushsubscriptionrequest
type PushSubscriptionRequest struct {
	XMLName         xml.Name `xml:"m:PushSubscriptionRequest"`
	FolderIds       FolderIds
	EventTypes      EventTypes
	Watermark       *string `xml:"t:Watermark,omitempty"`
	StatusFrequency *int    `xml:"t:StatusFrequency,omitempty"`
	URL             *string `xml:"t:URL"`
}

// The EventTypes element contains a collection of event notification
// types that are used to create a subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventtypes
type EventTypes struct {
	XMLName   xml.Name    `xml:"t:EventTypes"`
	EventType []EventType `xml:"t:EventType,omitempty"`
}

// The SendNotificationResult element contains the response of a client
// application to a push notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendnotificationresult
type SendNotificationResult struct {
	XMLName            xml.Name           `xml:"m:SendNotificationResult"`
	SubscriptionStatus SubscriptionStatus `xml:"m:SubscriptionStatus"`
}

type SubscribeResponseMessage struct {
	XMLName xml.Name `xml:"SubscribeResponseMessage"`
	Response
	SubscriptionId string `xml:"SubscriptionId"`
	Watermark      string `xml:"Watermark"`
}

type SendNotificationResponseMessage struct {
	Response
	Notification Notification
}

type Notification struct {
	XMLName           xml.Name            `xml:"Notification"`
	SubscriptionId    string              `xml:"SubscriptionId"`
	PreviousWatermark string              `xml:"PreviousWatermark"`
	MoreEvents        bool                `xml:"MoreEvents"`
	CreatedEvent      []NotificationEvent `xml:"CreatedEvent"`
	ModifiedEvent     []NotificationEvent `xml:"ModifiedEvent"`
	DeletedEvent      []NotificationEvent `xml:"DeletedEvent"`
}

type NotificationEvent struct {
	Watermark string    `xml:"Watermark"`
	TimeStamp time.Time `xml:"TimeStamp"`
	ItemId    ItemId    `xml:"ItemId"`
}
