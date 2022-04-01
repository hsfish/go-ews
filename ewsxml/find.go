package ewsxml

import (
	"encoding/xml"
	"time"
)

// Traversal defines whether the search finds items in folders or the folders'
// dumpsters.
type Traversal string

// The BaseShape element identifies the set of properties to return in an item
// or folder response.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseshape
type BaseShape string

// The BodyType element identifies how the body text is formatted in the
// response.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodytype
type BodyType string

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// Traversal_Shallow returns only the identities of items in the folder.
	Traversal_Shallow Traversal = "Shallow"
	// Traversal_SoftDeleted returns only the identities of items that are in a
	// folder's dumpster. Note that a soft-deleted traversal combined with a
	// search restriction will result in zero items returned even if there are
	// items that match the search criteria.
	Traversal_SoftDeleted Traversal = "SoftDeleted"
	// Traversal_Associated returns only the identities of associated items in
	// the folder.
	Traversal_Associated Traversal = "Associated"

	// BaseShape_IdOnly returns only the item or folder ID.
	BaseShape_IdOnly BaseShape = "IdOnly"
	// BaseShape_Default returns a set of properties that are defined as the
	// default for the item or folder.
	BaseShape_Default BaseShape = "Default"
	// BaseShape_AllProperties returns all the properties used by the Exchange
	// Business Logic layer to construct a folder.
	BaseShape_AllProperties BaseShape = "AllProperties"

	// BodyType_Best indicates the response will return the richest available
	// content of body text. This is useful if it is unknown whether the content
	// is text or HTML.
	// The returned body will be text if the stored body is plain text.
	// Otherwise, the response will return HTML if the stored body is in either
	// HTML or RTF format.
	BodyType_Best BodyType = "Best"
	// BodyType_HTML indicates the response will return an item body as HTML.
	BodyType_HTML BodyType = "HTML"
	// BodyType_Text indicates the response will return an item body as plain
	// text.
	BodyType_Text BodyType = "Text"
)

// The Body element specifies the body of an item.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/body
type Body struct {
	XMLName     xml.Name `xml:"t:Body"`
	BodyType    BodyType `xml:"BodyType,attr"`
	IsTruncated bool     `xml:"IsTruncated,attr"`
	Contents    []byte   `xml:",chardata"`
}

// The FindItem element defines a request to find items in a mailbox.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditem
type FindItem struct {
	XMLName   xml.Name  `xml:"m:FindItem"`
	Traversal Traversal `xml:"Traversal,attr"`
	ItemShape ItemShape
}

// The ItemShape element identifies a set of properties to return in a GetItem
// operation, FindItem operation, or SyncFolderItems operation response.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemshape
type ItemShape struct {
	XMLName            xml.Name  `xml:"m:ItemShape"`
	BaseShape          BaseShape `xml:"t:BaseShape,omitempty"`
	IncludeMimeContent bool      `xml:"t:IncludeMimeContent,omitempty"`
	BodyType           BodyType  `xml:"t:BodyType,omitempty"`
	FilterHtmlContent  bool      `xml:"t:FilterHtmlContent,omitempty"`
	// ConvertHtmlCodePageToUTF8
	// AdditionalProperties
}

type IndexedPageItemView struct {
	XMLName xml.Name `xml:"IndexedPageItemView"`
}

type FractionalPageItemView struct {
	XMLName xml.Name `xml:"FractionalPageItemView"`
}

type CalendarView struct {
	XMLName            xml.Name  `xml:"m:CalendarView"`
	MaxEntriesReturned uint      `xml:"MaxEntriesReturned,attr,omitempty"`
	StartDate          time.Time `xml:"StartDate,attr"`
	EndDate            time.Time `xml:"EndDate,attr"`
}

type ContactsView struct {
	XMLName xml.Name `xml:"ContactsView"`
}

// The FindItemResponseMessage element contains the status and result of a
// single FindItem operation request.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditemresponsemessage
type FindItemResponseMessage struct {
	Response
	XMLName            xml.Name `xml:"m:FindItemResponseMessage"`
	DescriptiveLinkKey int      `xml:"DescriptiveLinkKey,attr"`
	RootFolder         RootFolder
}

// The RootFolder element contains the results of a search of a single root
// folder during a FindItem operation.
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/rootfolder-finditemresponsemessage
type RootFolder struct {
	XMLName                 xml.Name `xml:"m:RootFolder"`
	IndexedPagingOffset     int      `xml:"IndexedPagingOffset,attr"`
	NumeratorOffset         int      `xml:"NumeratorOffset,attr"`
	AbsoluteDenominator     int      `xml:"AbsoluteDenominator,attr"`
	IncludesLastItemInRange bool     `xml:"IncludesLastItemInRange,attr"`
	TotalItemsInView        int      `xml:"TotalItemsInView,attr"`
	Items                   Items    `xml:"t:Items"`
	// Groups
}
