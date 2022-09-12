package ewsxml

import (
	"encoding/xml"
	"time"
)

type Recurrence struct {
	XMLName xml.Name `xml:"Recurrence"`

	DailyRecurrence *DailyRecurrence

	WeeklyRecurrence *WeeklyRecurrence

	AbsoluteMonthlyRecurrence *AbsoluteMonthlyRecurrence
	RelativeMonthlyRecurrence *RelativeMonthlyRecurrence

	AbsoluteYearlyRecurrence *AbsoluteYearlyRecurrence
	RelativeYearlyRecurrence *RelativeYearlyRecurrence

	EndDateRecurrence  *EndDateRecurrence
	NoEndRecurrence    *NoEndRecurrence
	NumberedRecurrence *NumberedRecurrence
}

type DailyRecurrence struct {
	XMLName  xml.Name `xml:"DailyRecurrence"`
	Interval int      `xml:"Interval"`
}

type WeeklyRecurrence struct {
	XMLName        xml.Name `xml:"WeeklyRecurrence"`
	Interval       int      `xml:"Interval"`
	DaysOfWeek     string   `xml:"DaysOfWeek,omitempty"`
	FirstDayOfWeek string   `xml:"FirstDayOfWeek,omitempty"`
}

type AbsoluteMonthlyRecurrence struct {
	XMLName    xml.Name `xml:"AbsoluteMonthlyRecurrence"`
	DayOfMonth int      `xml:"DayOfMonth,omitempty"`
	Interval   int      `xml:"Interval,omitempty"`
}

type RelativeMonthlyRecurrence struct {
	XMLName        xml.Name `xml:"RelativeMonthlyRecurrence"`
	DaysOfWeek     string   `xml:"DaysOfWeek,omitempty"`
	DayOfWeekIndex string   `xml:"DayOfWeekIndex"`
	Interval       int      `xml:"Interval,omitempty"`
}

type AbsoluteYearlyRecurrence struct {
	XMLName    xml.Name `xml:"AbsoluteYearlyRecurrence"`
	DayOfMonth int      `xml:"DayOfMonth,omitempty"`
	Month      string   `xml:"Month,omitempty"`
}

type RelativeYearlyRecurrence struct {
	XMLName        xml.Name `xml:"RelativeYearlyRecurrence"`
	DaysOfWeek     string   `xml:"DaysOfWeek,omitempty"`
	DayOfWeekIndex string   `xml:"DayOfWeekIndex"`
	Month          string   `xml:"Month,omitempty"`
}

type EndDateRecurrence struct {
	XMLName   xml.Name `xml:"EndDateRecurrence"`
	StartDate string   `xml:"StartDate"`
	EndDate   string   `xml:"EndDate"`
}

type NoEndRecurrence struct {
	XMLName   xml.Name `xml:"NoEndRecurrence"`
	StartDate string   `xml:"StartDate"`
}

type NumberedRecurrence struct {
	XMLName             xml.Name `xml:"NumberedRecurrence"`
	StartDate           string   `xml:"StartDate"`
	NumberOfOccurrences int      `xml:"NumberOfOccurrences"`
}

type SendRecurrence struct {
	XMLName xml.Name `xml:"t:Recurrence"`

	DailyRecurrence *SendDailyRecurrence

	WeeklyRecurrence *SendWeeklyRecurrence

	AbsoluteMonthlyRecurrence *SendAbsoluteMonthlyRecurrence
	RelativeMonthlyRecurrence *SendRelativeMonthlyRecurrence

	AbsoluteYearlyRecurrence *SendAbsoluteYearlyRecurrence
	RelativeYearlyRecurrence *SendRelativeYearlyRecurrence

	EndDateRecurrence  *SendEndDateRecurrence
	NoEndRecurrence    *SendNoEndRecurrence
	NumberedRecurrence *SendNumberedRecurrence
}

type SendDailyRecurrence struct {
	XMLName  xml.Name `xml:"t:DailyRecurrence"`
	Interval int      `xml:"t:Interval"`
}

type SendWeeklyRecurrence struct {
	XMLName        xml.Name `xml:"t:WeeklyRecurrence"`
	Interval       int      `xml:"t:Interval"`
	DaysOfWeek     string   `xml:"t:DaysOfWeek,omitempty"`
	FirstDayOfWeek string   `xml:"t:FirstDayOfWeek,omitempty"`
}

type SendAbsoluteMonthlyRecurrence struct {
	XMLName    xml.Name `xml:"t:AbsoluteMonthlyRecurrence"`
	DayOfMonth int      `xml:"t:DayOfMonth,omitempty"`
	Interval   int      `xml:"t:Interval,omitempty"`
}

type SendRelativeMonthlyRecurrence struct {
	XMLName        xml.Name `xml:"t:RelativeMonthlyRecurrence"`
	DaysOfWeek     string   `xml:"t:DaysOfWeek,omitempty"`
	DayOfWeekIndex string   `xml:"t:DayOfWeekIndex"`
	Interval       int      `xml:"t:Interval,omitempty"`
}

type SendAbsoluteYearlyRecurrence struct {
	XMLName    xml.Name `xml:"t:AbsoluteYearlyRecurrence"`
	DayOfMonth int      `xml:"t:DayOfMonth,omitempty"`
	Month      string   `xml:"t:Month,omitempty"`
}

type SendRelativeYearlyRecurrence struct {
	XMLName        xml.Name `xml:"t:RelativeYearlyRecurrence"`
	DaysOfWeek     string   `xml:"t:DaysOfWeek,omitempty"`
	DayOfWeekIndex string   `xml:"t:DayOfWeekIndex"`
	Month          string   `xml:"t:Month,omitempty"`
}

type SendEndDateRecurrence struct {
	XMLName   xml.Name  `xml:"t:EndDateRecurrence"`
	StartDate time.Time `xml:"t:StartDate"`
	EndDate   time.Time `xml:"t:EndDate"`
}

type SendNoEndRecurrence struct {
	XMLName   xml.Name  `xml:"t:NoEndRecurrence"`
	StartDate time.Time `xml:"t:StartDate"`
}

type SendNumberedRecurrence struct {
	XMLName             xml.Name  `xml:"t:NumberedRecurrence"`
	StartDate           time.Time `xml:"t:StartDate"`
	NumberOfOccurrences int       `xml:"t:NumberOfOccurrences"`
}
