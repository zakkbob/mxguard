package mail

import (
	"maps"
)

type Attachment struct {
	fileName    string
	path        string
	contentType string
}

type MailInfo struct {
	From        string
	To          []string
	Subject     string 
	Bcc         []string
	Cc          []string
	ReplyTo     string
	Html        string
	Text        string
	Attachments []*Attachment
	Headers     map[string]string
}
