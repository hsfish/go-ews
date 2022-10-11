package ewsxml

import (
	"encoding/xml"
	"fmt"
	"time"
)

type TimeWindow struct {
	StartTime time.Time `xml:"t:StartTime"`
	EndTime   time.Time `xml:"t:EndTime"`
}

type TimeWindowStr struct {
	StartTime string `xml:"t:StartTime"`
	EndTime   string `xml:"t:EndTime"`
}

type TimeZone struct {
	Bias         int          `xml:"t:Bias"`
	StandardTime TimeZoneTime `xml:"t:StandardTime"`
	DaylightTime TimeZoneTime `xml:"t:DaylightTime"`
}

type TimeZoneTime struct {
	Bias      int    `xml:"t:Bias"`
	Time      string `xml:"t:Time"`
	DayOrder  int16  `xml:"t:DayOrder"`
	Month     int16  `xml:"t:Month"`
	DayOfWeek string `xml:"t:DayOfWeek"`
	Year      string `xml:"Year,omitempty"`
}

type TimeZoneContext struct {
	XMLName            xml.Name `xml:"t:TimeZoneContext"`
	TimeZoneDefinition *TimeZoneDefinition
}

type TimeZoneDefinition struct {
	XMLName xml.Name `xml:"t:TimeZoneDefinition"`
	Id      string   `xml:"Id,attr,omitempty"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttimezone
type StartTimeZone struct {
	XMLName xml.Name `xml:"t:StartTimeZone"`
	Id      string   `xml:"Id,attr,omitempty"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/endtimezone
type EndTimeZone struct {
	XMLName xml.Name `xml:"t:EndTimeZone"`
	Id      string   `xml:"Id,attr,omitempty"`
}

type Time string

func (t Time) ToTime() (time.Time, error) {
	offset, err := getRFC3339Offset(time.Now())
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, string(t)+offset)

}

// return RFC3339 formatted offset, ex: +03:00 -03:30
func getRFC3339Offset(t time.Time) (string, error) {

	_, offset := t.Zone()
	i := int(float32(offset) / 36)

	sign := "+"
	if i < 0 {
		i = -i
		sign = "-"
	}
	hour := i / 100
	min := i % 100
	min = (60 * min) / 100

	return fmt.Sprintf("%s%02d:%02d", sign, hour, min), nil
}
