package ews

import (
	"context"
	"encoding/xml"

	"github.com/soft-stech/go-ews/ewsxml"
)

type GetRoomListsOperation struct {
	header ewsxml.Header
	body   struct {
		XMLName xml.Name `xml:"m:GetRoomLists"`
	}
}

func (op *GetRoomListsOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetRoomListsOperation) Body() interface{}      { return op.body }

type GetRoomListsResponse struct {
	ewsxml.Response
	RoomLists struct {
		Address []ewsxml.Mailbox `xml:"Address"`
	} `xml:"RoomLists"`
}

func GetRoomLists(ctx context.Context, req Requester, op *GetRoomListsOperation) (*GetRoomListsResponse, error) {
	var out GetRoomListsResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getrooms-operation
type GetRoomsOperation struct {
	header   ewsxml.Header
	GetRooms ewsxml.GetRooms
}

func (op *GetRoomsOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetRoomsOperation) Body() interface{}      { return op.GetRooms }

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

func GetRooms(ctx context.Context, req Requester, op *GetRoomsOperation) (*GetRoomsResponse, error) {
	var out GetRoomsResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
