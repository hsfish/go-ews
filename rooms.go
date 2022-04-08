package ews

import (
	"context"
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

func GetRoomLists(ctx context.Context, req Requester) (*GetRoomListsResponse, error) {
	var out GetRoomListsResponse
	return &out, requestAndUnmarshal(ctx, req, GetRoomListsRequest{}, &out)
}

type GetRoomsRequest struct {
	XMLName  xml.Name `xml:"m:GetRooms"`
	RoomList struct {
		EmailAddress string `xml:"t:EmailAddress"`
	} `xml:"m:RoomList"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getroomsresponse
type GetRoomsResponse struct {
	ewsxml.Response
	Rooms []struct {
		Name         string
		EmailAddress string
		RoutingType  ewsxml.RoutingType
		MailboxType  ewsxml.MailboxType
	} `xml:"Rooms>Room>Id"`
}

func GetRooms(ctx context.Context, req Requester, email string) (*GetRoomsResponse, error) {
	var in GetRoomsRequest
	in.RoomList.EmailAddress = email

	var out GetRoomsResponse
	return &out, requestAndUnmarshal(ctx, req, in, &out)
}
