# internal/mail

for receiving, processing, and sending emails

## Receive
a wrapper over go-smtp which pushes received mail into a channel


Features:
- [ ] can receive emails
- [ ] handles attachments
- [ ] carries out dkim, spf, and dmarc checks
- [ ] enforce STARTTLS