package ewsxml

import "encoding/xml"

// The UserConfigurationProperties element specifies the property types to get
// in a GetUserConfiguration operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userconfigurationproperties
type UserConfigurationProperties string

func (s UserConfigurationProperties) String() string { return string(s) }

const (
	UserConfigurationProperties_Id         UserConfigurationProperties = "Id"
	UserConfigurationProperties_Dictionary UserConfigurationProperties = "Dictionary"
	UserConfigurationProperties_XmlData    UserConfigurationProperties = "XmlData"
	UserConfigurationProperties_BinaryData UserConfigurationProperties = "BinaryData"
	UserConfigurationProperties_All        UserConfigurationProperties = "All"
)

// The GetUserConfiguration element represent a request to get a user
// configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserconfiguration
type GetUserConfiguration struct {
	XMLName                     xml.Name `xml:"m:GetUserConfiguration"`
	UserConfigurationName       SendUserConfigurationName
	UserConfigurationProperties UserConfigurationProperties `xml:"m:UserConfigurationProperties"`
}

type SendUserConfigurationName struct {
	XMLName               xml.Name `xml:"m:UserConfigurationName"`
	Name                  string   `xml:"Name,attr"`
	DistinguishedFolderId DistinguishedFolderId
}

type UserConfigurationName struct {
	XMLName xml.Name `xml:"UserConfigurationName"`
	Name    string   `xml:"Name,attr"`
	//DistinguishedFolderId DistinguishedFolderId
}

type GetUserConfigurationResponseMessage struct {
	UserConfiguration UserConfiguration
}

type UserConfiguration struct {
	XMLName               xml.Name              `xml:"UserConfiguration"`
	UserConfigurationName UserConfigurationName `xml:"UserConfigurationName"`
	ItemId                ItemId
	XmlData               string
	//Dictionary
	//BinaryData
}

type UpdateUserConfiguration struct {
	XMLName           xml.Name `xml:"m:UpdateUserConfiguration"`
	UserConfiguration SendUserConfiguration
}

type SendUserConfiguration struct {
	XMLName               xml.Name `xml:"m:UserConfiguration"`
	UserConfigurationName struct {
		Name                  string `xml:"Name,attr"`
		DistinguishedFolderId DistinguishedFolderId
	} `xml:"t:UserConfigurationName"`
	XmlData *string `xml:"t:XmlData,omitempty"`
	//Dictionary
	//BinaryData
	//Dictionary
}
