package ewsxml

import (
	"encoding/xml"
)

// The UpdateFolder element represents the operation that is used to update
// properties for a specified folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolder
type UpdateFolder struct {
	XMLName       xml.Name `xml:"m:UpdateFolder"`
	FolderChanges FolderChanges
}

// The FolderChanges element represents a collection of changes for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderchanges
type FolderChanges struct {
	XMLName      xml.Name `xml:"m:FolderChanges"`
	FolderChange []FolderChange
}

// The FolderChange element represents a collection of changes to be
// performed on a single folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderchange
type FolderChange struct {
	XMLName               xml.Name `xml:"t:FolderChange"`
	FolderId              *SendFolderId
	DistinguishedFolderId *SendDistinguishedFolderId
	Updates               FolderUpdates
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updates-folder
type FolderUpdates struct {
	XMLName xml.Name `xml:"t:Updates"`
	//AppendToFolderField []AppendToFolderField
	SetFolderField []SetFolderField
	//DeleteFolderField []DeleteFolderField
}

// The UpdateFolderResponseMessage element contains the status and result
// of updates defined by the FolderChange element of an UpdateFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolderresponsemessage
type UpdateFolderResponseMessage struct {
	XMLName xml.Name `xml:"UpdateFolderResponseMessage"`
	Response
	Folders Folders
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setfolderfield
type SetFolderField struct {
	XMLName        xml.Name `xml:"t:SetFolderField"`
	FieldURI       FieldURI
	CalendarFolder *SendCalendarFolder
}
