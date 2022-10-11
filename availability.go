package ews

import (
	"context"
	"encoding/xml"

	"github.com/soft-stech/go-ews/ewsxml"
)

type GetUserAvailabilityOperation struct {
	header                     ewsxml.Header
	GetUserAvailabilityRequest ewsxml.GetUserAvailabilityRequest
}

func (op *GetUserAvailabilityOperation) Header() *ewsxml.Header { return &op.header }
func (op *GetUserAvailabilityOperation) Body() interface{}      { return op.GetUserAvailabilityRequest }

type GetUserAvailabilityResponse struct {
	XMLName               xml.Name `xml:"GetUserAvailabilityResponse"`
	FreeBusyResponseArray ewsxml.FreeBusyResponseArray
}

func GetUserAvailability(ctx context.Context, req Requester, op *GetUserAvailabilityOperation) (*GetUserAvailabilityResponse, error) {
	if op.GetUserAvailabilityRequest.FreeBusyViewOptions.RequestedView == "" {
		op.GetUserAvailabilityRequest.FreeBusyViewOptions.RequestedView = ewsxml.RequestedView_Detailed
	}

	if op.GetUserAvailabilityRequest.FreeBusyViewOptions.MergedFreeBusyIntervalInMinutes == 0 {
		op.GetUserAvailabilityRequest.FreeBusyViewOptions.MergedFreeBusyIntervalInMinutes = 30
	}
	if op.GetUserAvailabilityRequest.SuggestionsViewOptions == nil {
		op.GetUserAvailabilityRequest.SuggestionsViewOptions = &ewsxml.SuggestionsViewOptions{
			GoodThreshold:                  49,
			MaximumResultsByDay:            2,
			MaximumNonWorkHourResultsByDay: 0,
			MeetingDurationInMinutes:       60,
			MinimumSuggestionQuality:       ewsxml.MinimumSuggestionQuality_Good,
			DetailedSuggestionsWindow: ewsxml.TimeWindowStr{
				StartTime: op.GetUserAvailabilityRequest.FreeBusyViewOptions.TimeWindow.StartTime,
				EndTime:   op.GetUserAvailabilityRequest.FreeBusyViewOptions.TimeWindow.EndTime,
			},
		}
	}
	var out GetUserAvailabilityResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
