package main

import (
	"log"

	"github.com/zakkbob/mxguard/internal/mail"
)

func main() {
	s := mail.NewServer("localhost:1025", "localhost")
	go s.ListenAndServe()

	for {
		Mail := <-s.Mail
		log.Printf("%w", Mail)
	}
}
