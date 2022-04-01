package ewsxml

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/message-ex15websvcsotherref
type Message struct {
	// MimeContent string `xml:"t:MimeContent"`
	// ItemId         ItemId         `xml:"t:ItemId"`
	// ParentFolderId ParentFolderId `xml:"t:ParentFolderId"`
	ItemClass   string      `xml:"t:ItemClass"`
	Subject     string      `xml:"t:Subject"`
	Sensitivity Sensitivity `xml:"t:Sensitivity"`
	Body        Body        `xml:"t:Body"`
	// Attachments                  string      `xml:"t:Attachments"`
	// DateTimeReceived             string      `xml:"t:DateTimeReceived"`
	// Size                         string      `xml:"t:Size"`
	// Categories                   string      `xml:"t:Categories"`
	// Importance                   string      `xml:"t:Importance"`
	// InReplyTo                    string      `xml:"t:InReplyTo"`
	// IsSubmitted                  string      `xml:"t:IsSubmitted"`
	// IsDraft                      string      `xml:"t:IsDraft"`
	// IsFromMe                     string      `xml:"t:IsFromMe"`
	// IsResend                     string      `xml:"t:IsResend"`
	// IsUnmodified                 string      `xml:"t:IsUnmodified"`
	// InternetMessageHeaders       string      `xml:"t:InternetMessageHeaders"`
	// DateTimeSent                 string      `xml:"t:DateTimeSent"`
	// DateTimeCreated              string      `xml:"t:DateTimeCreated"`
	// ResponseObjects              string      `xml:"t:ResponseObjects"`
	// ReminderDueBy                string      `xml:"t:ReminderDueBy"`
	// ReminderIsSet                string      `xml:"t:ReminderIsSet"`
	// ReminderMinutesBeforeStart   string      `xml:"t:ReminderMinutesBeforeStart"`
	// DisplayCc                    string      `xml:"t:DisplayCc"`
	// DisplayTo                    string      `xml:"t:DisplayTo"`
	// HasAttachments               string      `xml:"t:HasAttachments"`
	// ExtendedProperty             string      `xml:"t:ExtendedProperty"`
	// Culture                      string      `xml:"t:Culture"`
	Sender       OneMailbox `xml:"t:Sender"`
	ToRecipients []Mailbox  `xml:"t:ToRecipients"`
	// CcRecipients                 string      `xml:"t:CcRecipients"`
	// BccRecipients                string      `xml:"t:BccRecipients"`
	// IsReadReceiptRequested       string      `xml:"t:IsReadReceiptRequested"`
	// IsDeliveryReceiptRequested   string      `xml:"t:IsDeliveryReceiptRequested"`
	// ConversationIndex            string      `xml:"t:ConversationIndex"`
	// ConversationTopic            string      `xml:"t:ConversationTopic"`
	// From                         string      `xml:"t:From"`
	// InternetMessageId            string      `xml:"t:InternetMessageId"`
	// IsRead                       string      `xml:"t:IsRead"`
	// IsResponseRequested          string      `xml:"t:IsResponseRequested"`
	// References                   string      `xml:"t:References"`
	// ReplyTo                      string      `xml:"t:ReplyTo"`
	// EffectiveRights              string      `xml:"t:EffectiveRights"`
	// ReceivedBy                   string      `xml:"t:ReceivedBy"`
	// ReceivedRepresenting         string      `xml:"t:ReceivedRepresenting"`
	// LastModifiedName             string      `xml:"t:LastModifiedName"`
	// LastModifiedTime             string      `xml:"t:LastModifiedTime"`
	// IsAssociated                 string      `xml:"t:IsAssociated"`
	// WebClientReadFormQueryString string      `xml:"t:WebClientReadFormQueryString"`
	// WebClientEditFormQueryString string      `xml:"t:WebClientEditFormQueryString"`
	// ConversationId               string      `xml:"t:ConversationId"`
	// UniqueBody                   string      `xml:"t:UniqueBody"`
	// ReminderMessageData          string      `xml:"t:ReminderMessageData"`
}
