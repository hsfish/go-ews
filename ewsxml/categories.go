package ewsxml

import "encoding/xml"

type Categories struct {
	XMLName xml.Name `xml:"Categories"`
	String  []string `xml:"String,omitempty"`
}

type SendCategories struct {
	XMLName xml.Name `xml:"t:Categories"`
	String  []string `xml:"t:String,omitempty"`
}
