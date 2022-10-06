package ewsxml

import (
	"encoding/xml"
)

// The FolderIds element contains an array of folder identifiers
// that are used to identify folders to copy, move, get, delete,
// or monitor for event notifications.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderids
type FolderIds struct {
	XMLName               xml.Name `xml:"t:FolderIds"`
	FolderId              []FolderId
	DistinguishedFolderId []DistinguishedFolderId
}

// The FolderId element contains the identifier and change key of a folder
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderid
type FolderId struct {
	XMLName   xml.Name `xml:"t:FolderId"`
	Id        string   `xml:"Id,attr"`
	ChangeKey string   `xml:"t:ChangeKey,attr,omitempty"`
}

// The DistinguishedFolderId element identifies folders that can be referenced
// by name. If you do not use this element, you must use the FolderId element to
// identify a folder.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
type DistinguishedFolderId struct {
	XMLName xml.Name `xml:"t:DistinguishedFolderId"`
	Id      string   `xml:"Id,attr,omitempty"`
	Mailbox Mailbox  `xml:"t:Mailbox"`
}

// The ParentFolderId element represents the identifier of the parent folder
// that contains the item or folder.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderid
type ParentFolderId struct {
	XMLName   xml.Name `xml:"m:ParentFolderId"`
	Id        string   `xml:"Id,attr"`
	ChangeKey string   `xml:"ChangeKey,attr,omitempty"`
}
