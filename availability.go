package ews

import (
	"context"
	"encoding/xml"

	"github.com/hsfish/go-ews/ewsxml"
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
	SuggestionsResponse   ewsxml.SuggestionsResponse
}

func GetUserAvailability(ctx context.Context, req Requester, op *GetUserAvailabilityOperation) (*GetUserAvailabilityResponse, error) {
	if op.GetUserAvailabilityRequest.FreeBusyViewOptions.RequestedView == "" {
		op.GetUserAvailabilityRequest.FreeBusyViewOptions.RequestedView = ewsxml.RequestedView_Detailed
	}

	if op.GetUserAvailabilityRequest.FreeBusyViewOptions.MergedFreeBusyIntervalInMinutes == 0 {
		op.GetUserAvailabilityRequest.FreeBusyViewOptions.MergedFreeBusyIntervalInMinutes = 30
	}

	if op.GetUserAvailabilityRequest.SuggestionsViewOptions != nil {
		if op.GetUserAvailabilityRequest.SuggestionsViewOptions.GoodThreshold == 0 {
			op.GetUserAvailabilityRequest.SuggestionsViewOptions.GoodThreshold = 49
		}

		if op.GetUserAvailabilityRequest.SuggestionsViewOptions.MaximumResultsByDay == 0 {
			op.GetUserAvailabilityRequest.SuggestionsViewOptions.MaximumResultsByDay = 2
		}
		if op.GetUserAvailabilityRequest.SuggestionsViewOptions.MeetingDurationInMinutes == 0 {
			op.GetUserAvailabilityRequest.SuggestionsViewOptions.MeetingDurationInMinutes = 60

		}
		if op.GetUserAvailabilityRequest.SuggestionsViewOptions.MinimumSuggestionQuality == "" {
			op.GetUserAvailabilityRequest.SuggestionsViewOptions.MinimumSuggestionQuality = ewsxml.MinimumSuggestionQuality_Good

		}

		if op.GetUserAvailabilityRequest.SuggestionsViewOptions.DetailedSuggestionsWindow.StartTime == "" {
			t := op.GetUserAvailabilityRequest.FreeBusyViewOptions.TimeWindow.StartTime
			op.GetUserAvailabilityRequest.SuggestionsViewOptions.DetailedSuggestionsWindow.StartTime = t
		}

		if op.GetUserAvailabilityRequest.SuggestionsViewOptions.DetailedSuggestionsWindow.EndTime == "" {
			t := op.GetUserAvailabilityRequest.FreeBusyViewOptions.TimeWindow.EndTime
			op.GetUserAvailabilityRequest.SuggestionsViewOptions.DetailedSuggestionsWindow.EndTime = t
		}

	}

	var out GetUserAvailabilityResponse
	return &out, req.Request(NewOperationRequest(ctx, op), &out)
}
