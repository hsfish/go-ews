package ewsxml

import (
	"encoding/xml"
)

type GetItem struct {
	XMLName   xml.Name `xml:"m:GetItem"`
	ItemShape ItemShape
	ItemsIds  *ItemIds
}

// The GetItemResponseMessage element contains the status and result of a
// single GetItem operation request.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitemresponsemessage
type GetItemResponseMessage struct {
	XMLName xml.Name `xml:"GetItemResponseMessage"`
	Response
	Items Items
}
