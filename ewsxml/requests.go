package ewsxml

import (
	"encoding/xml"
)

type GetRooms struct {
	XMLName  xml.Name `xml:"m:GetRooms"`
	RoomList struct {
		EmailAddress string `xml:"t:EmailAddress"`
	} `xml:"m:RoomList"`
}
