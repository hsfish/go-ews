package ews

import (
	"encoding/xml"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type GetRoomListsRequest struct {
	XMLName xml.Name `xml:"m:GetRoomLists"`
}

type GetRoomListsResponse struct {
	ewsxml.Response
	RoomLists struct {
		Address []ewsxml.Mailbox `xml:"Address"`
	} `xml:"RoomLists"`
}

func GetRoomLists(req Requester) (*GetRoomListsResponse, error) {
	var out GetRoomListsResponse
	return &out, requestAndUnmarshal(req, GetRoomListsRequest{}, &out)
}

type GetRoomsRequest struct {
	XMLName  xml.Name `xml:"m:GetRooms"`
	RoomList struct {
		EmailAddress string `xml:"t:EmailAddress"`
	} `xml:"m:RoomList"`
}

type GetRoomsResponse struct {
	ewsxml.Response
	Rooms struct {
		Room []struct {
			Id ewsxml.Mailbox `xml:"Id"`
		} `xml:"Room"`
	} `xml:"Rooms"`
}

func GetRooms(req Requester, email string) (*GetRoomsResponse, error) {
	var in GetRoomsRequest
	in.RoomList.EmailAddress = email

	var out GetRoomsResponse
	return &out, requestAndUnmarshal(req, in, &out)
}
