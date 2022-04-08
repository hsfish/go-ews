package ewsxml

type BasePoint string

func (s BasePoint) String() string { return string(s) }

const (
	BasePointBeginning BasePoint = "Beginning"
	BasePointEnd       BasePoint = "End"
)

type Persona struct {
	PersonaId            PersonaId            `xml:"PersonaId"`
	DisplayName          string               `xml:"DisplayName"`
	Title                string               `xml:"Title"`
	Department           string               `xml:"Department"`
	Departments          Departments          `xml:"Departments"`
	EmailAddress         Mailbox              `xml:"EmailAddress"`
	RelevanceScore       int                  `xml:"RelevanceScore"`
	BusinessPhoneNumbers BusinessPhoneNumbers `xml:"BusinessPhoneNumbers"`
	MobilePhones         MobilePhones         `xml:"MobilePhones"`
	OfficeLocations      OfficeLocations      `xml:"OfficeLocations"`
}

type PersonaId struct {
	Id string `xml:"Id,attr"`
}

type BusinessPhoneNumbers struct {
	PhoneNumberAttributedValue PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

type MobilePhones struct {
	PhoneNumberAttributedValue PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

type Value struct {
	Number string `json:"Number"`
	Type   string `json:"Type"`
}

type PhoneNumberAttributedValue struct {
	Value Value `json:"Value"`
}

type OfficeLocations struct {
	StringAttributedValue StringAttributedValue `xml:"StringAttributedValue"`
}

type Departments struct {
	StringAttributedValue StringAttributedValue `xml:"StringAttributedValue"`
}

type StringAttributedValue struct {
	Value string `json:"Value"`
}
