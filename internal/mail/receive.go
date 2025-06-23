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

// A Session is returned after successful login.
type Session struct {
	from    string
	to      []string
	backend *Backend
}

type Server struct {
	server *smtp.Server
	Mail   chan MailInfo
}

// Called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{
		backend: bkd,
	}, nil
}

// AuthMechanisms returns a slice of available auth mechanisms;
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

// Called when 'MAIL FROM:' command is given
func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	log.Println("Mail from:", from)
	return nil
}

// Called when 'RCPT TO:' command is given
func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.to = append(s.to, to)
	log.Println("Rcpt to:", to)
	return nil
}

// Called when 'DATA' command is given
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
