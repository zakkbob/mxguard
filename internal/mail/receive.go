package mail

import (
	"io"
	"log"
	"time"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/jhillyerd/enmime"
)

// The Backend implements SMTP server methods.
type Backend struct {
	Mail chan MailInfo
}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{
		backend: bkd,
	}, nil
}

// A Session is returned after successful login.
type Session struct {
	from    string
	to      []string
	backend *Backend
}

// AuthMechanisms returns a slice of available auth mechanisms; only PLAIN is
// supported in this example.
func (s *Session) AuthMechanisms() []string {
	return []string{sasl.Anonymous}
}

// Auth is the handler for supported authenticators.
func (s *Session) Auth(mech string) (sasl.Server, error) {
	return sasl.NewAnonymousServer(func(trace string) error {
		log.Printf("New anonymous login with trace: %s", trace) // Note - does this log need to exist
		return nil
	}), nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.to = append(s.to, to)
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if e, err := enmime.ReadEnvelope(r); err != nil {
		return err
	} else {
		m := MailInfo{
			EnvelopeFrom: s.from,
			EnvelopeTo:   s.to,
			ReplyTo:      s.from,
			From:         e.GetHeader("From"),
			To:           e.GetHeaderValues("To"),
			Subject:      e.GetHeader("Subject"),
			Html:         e.HTML,
			Text:         e.Text,
		}

		s.backend.Mail <- m
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

type Server struct {
	server *smtp.Server
	Mail   chan MailInfo
}

func NewServer(addr, domain string) *Server {
	b := &Backend{
		Mail: make(chan MailInfo), // NOTE - make buffered?
	}

	s := smtp.NewServer(b)

	s.Addr = addr
	s.Domain = domain
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = false

	return &Server{
		server: s,
		Mail:   b.Mail,
	}
}

func (s *Server) ListenAndServe() error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

// ExampleServer runs an example SMTP server.
//
// It can be tested manually with e.g. netcat:
//
//	> netcat -C localhost 1025
//	EHLO localhost
//	AUTH PLAIN
//	AHVzZXJuYW1lAHBhc3N3b3Jk
//	MAIL FROM:<root@nsa.gov>
//	RCPT TO:<root@gchq.gov.uk>
//	DATA
//	Hey <3
//	.
