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

type SendFolderIds struct {
	XMLName               xml.Name `xml:"m:FolderIds"`
	FolderId              []SendFolderId
	DistinguishedFolderId []SendDistinguishedFolderId
}

// The ToFolderId element represents the destination folder for a copied or moved item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tofolderid
type ToFolderId struct {
	XMLName               xml.Name `xml:"m:ToFolderId"`
	FolderId              []SendFolderId
	DistinguishedFolderId []SendDistinguishedFolderId
}

type SendParentFolderId struct {
	XMLName               xml.Name `xml:"m:ParentFolderId"`
	FolderId              *SendFolderId
	DistinguishedFolderId *SendDistinguishedFolderId
}
type SendParentFolderIds struct {
	XMLName               xml.Name `xml:"m:ParentFolderIds"`
	FolderId              []SendFolderId
	DistinguishedFolderId []SendDistinguishedFolderId
}

// The FolderId element contains the identifier and change key of a folder
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderid
type FolderId struct {
	XMLName   xml.Name `xml:"FolderId"`
	Id        string   `xml:"Id,attr"`
	ChangeKey string   `xml:"ChangeKey,attr,omitempty"`
}

type SendFolderId struct {
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

type SendDistinguishedFolderId struct {
	XMLName xml.Name    `xml:"t:DistinguishedFolderId"`
	Id      string      `xml:"Id,attr,omitempty"`
	Mailbox SendMailbox `xml:"t:Mailbox"`
}

// The ParentFolderId element represents the identifier of the parent folder
// that contains the item or folder.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderid
type ParentFolderId struct {
	XMLName   xml.Name `xml:"m:ParentFolderId"`
	Id        string   `xml:"Id,attr"`
	ChangeKey string   `xml:"ChangeKey,attr,omitempty"`
}

// The GetFolder operation gets folders from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolder-operation
type GetFolder struct {
	XMLName     xml.Name `xml:"m:GetFolder"`
	FolderShape FolderShape
	FolderIds   SendFolderIds
}

// The FindFolder element defines a request to find folders in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findfolder
type FindFolder struct {
	XMLName     xml.Name  `xml:"m:FindFolder"`
	Traversal   Traversal `xml:"Traversal,attr"`
	FolderShape FolderShape
	//IndexedPageFolderView
	//FractionalPageFolderView
	//Restriction
	ParentFolderIds SendParentFolderIds
}

// The FolderShape element identifies the folder properties to include in a
// GetFolder, FindFolder, or SyncFolderHierarchy response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldershape
type FolderShape struct {
	XMLName   xml.Name  `xml:"m:FolderShape"`
	BaseShape BaseShape `xml:"t:BaseShape,omitempty"`
	// AdditionalProperties
}

// The GetFolderResponseMessage element contains the status and result
// of a single GetFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolderresponsemessage
type GetFolderResponseMessage struct {
	XMLName xml.Name `xml:"GetFolderResponseMessage"`
	Response
	Folders Folders
}

// The FindFolderResponseMessage element contains the status and result of a single
// FindFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findfolderresponsemessage
type FindFolderResponseMessage struct {
	XMLName xml.Name `xml:"FindFolderResponseMessage"`
	Response
	RootFolder struct {
		Folders Folders
	} `xml:"RootFolder"`
}

type Folders struct {
	XMLName xml.Name `xml:"Folders"`
	//Folder         []Folder
	CalendarFolder []CalendarFolder
	//ContactsFolder
	//TasksFolder
	//SearchFolder
}

type SendFolders struct {
	XMLName xml.Name `xml:"m:Folders"`
	//Folder         []Folder
	CalendarFolder []SendCalendarFolder
	//ContactsFolder
	//TasksFolder
	//SearchFolder
}

type CalendarFolder struct {
	XMLName          xml.Name `xml:"CalendarFolder"`
	FolderId         FolderId `xml:"FolderId"`
	ParentFolderId   ParentFolderId
	FolderClass      string
	DisplayName      string
	TotalCount       int
	ChildFolderCount int
	//ExtendedProperty
	//ManagedFolderInformation
	UnreadCount int
	//PermissionSet
	EffectiveRights struct {
		CreateAssociated bool
		CreateContents   bool
		CreateHierarchy  bool
		Delete           bool
		Modify           bool
		Read             bool
		ViewPrivateItems bool
	}
}

type SendCalendarFolder struct {
	XMLName     xml.Name `xml:"t:CalendarFolder"`
	DisplayName string   `xml:"t:DisplayName"`
}

// The CreateFolder operation creates folders, calendar folders, contacts folders,
// tasks folders, and search folders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolder-operation
type CreateFolder struct {
	XMLName        xml.Name `xml:"m:CreateFolder"`
	ParentFolderId *SendParentFolderId
	Folders        SendFolders
}

// The CreateFolderResponseMessage element contains the status
// and result of a single CreateFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderresponsemessage
type CreateFolderResponseMessage struct {
	XMLName xml.Name `xml:"CreateFolderResponseMessage"`
	Response
	Folders Folders
}
