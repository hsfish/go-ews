package ewsxml

import (
	"encoding/xml"
)

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/cancelcalendaritem
type SendCancelCalendarItem struct {
	XMLName         xml.Name `xml:"t:CancelCalendarItem"`
	ReferenceItemId ReferenceItemId
}
