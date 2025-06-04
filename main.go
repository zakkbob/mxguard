package main

import (
	"net/http"

	"github.com/zakkbob/mxguard/internal/auth"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("website")))
	http.HandleFunc("/api/register", auth.RegisterHandler)
	http.HandleFunc("/api/login", auth.LoginHandler)
	http.ListenAndServe(":8080", nil)
}
