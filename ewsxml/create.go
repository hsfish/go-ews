package ewsxml

import (
	"encoding/xml"
)

// The CreateItemResponseMessage element contains the status and result of a
// single CreateItem operation request.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitemresponsemessage
type CreateItemResponseMessage struct {
	XMLName xml.Name `xml:"CreateItemResponseMessage"`
	Response
	Items Items
}
