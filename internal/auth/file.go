// I hate this so much, but I have 30 mins :/
package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/zakkbob/mxguard/internal/utils"
)

const userFile = "users.txt" // Oh my, this is so bad

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	// does user exist already?
	users, _ := readUsers()
	if _, exists := users[req.Username]; exists {
		http.Error(w, "User exists", http.StatusBadRequest)
		return
	}

	hash, _ := utils.HashPassword(req.Password)
	f, _ := os.OpenFile(userFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(req.Username + ":" + string(hash) + "\n")

	w.WriteHeader(http.StatusCreated)
}

func readUsers() (map[string]string, error) {
	data, err := os.ReadFile(userFile)
	if err != nil {
		return map[string]string{}, nil
	}
	lines := strings.Split(string(data), "\n")
	users := map[string]string{}
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			users[parts[0]] = parts[1]
		}
	}
	return users, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	users, _ := readUsers()
	hash, exists := users[req.Username]
	if !exists || !utils.ValidatePasswordHash(req.Password, hash) {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: req.Username,
		Path:  "/",
	})
	w.WriteHeader(http.StatusOK)
}
