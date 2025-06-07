package mail

type Attachment struct {
	fileName    string
	path        string
	contentType string
}

type MailInfo struct {
	EnvelopeFrom string
	EnvelopeTo   []string
	From         string
	To           []string
	ReplyTo      string
	Subject      string
	Html         string
	Text         string
	Attachments  []*Attachment
}
