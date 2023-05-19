package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/yuiuae/HTTPServer/pkg/hasher"
)

type UserInfo struct {
	Passhash string
	Id       string
}

var usersTable = map[string]UserInfo{}

// Create a struct that models the structure for a user creating
// Request
type CrRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Response
type CrResponse struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

func handleUserCreate(w http.ResponseWriter, r *http.Request) {
	req := &CrRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	fmt.Println("Test = ", req.UserName)
	if err != nil || len(req.UserName) <= 4 || len(req.Password) <= 8 {
		// If there is something wrong with the request body, return a 400 status
		http.Error(w, "Bad request, empty username or password", 400)
		return
	}
	// Hash the password using the bcrypt algorithm
	hashedPassword, err := hasher.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Internal Server Error (hash)", 500)
		return
	}
	uid, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, "Internal Server Error (UUID)", 500)
		return
	}
	suid := fmt.Sprintf("%v", uid)
	resp := &CrResponse{suid, req.UserName}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error (json Encoder)", 500)
		return
	}

	usersTable[req.UserName] = UserInfo{hashedPassword, suid}

	// We reach this point if the credentials we correctly stored and the default status of 200 is sent back
}

type LogRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type LogResponse struct {
	URL string `json:"username"`
}

func handleUserLogin(w http.ResponseWriter, r *http.Request) {
	req := &LogRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", 400)
		return
	}
	fmt.Println(usersTable[req.UserName].Passhash, req.Password)
	ok := hasher.CheckPasswordHash(usersTable[req.UserName].Passhash, req.Password)
	if !ok {
		http.Error(w, "Invalid username/password", 400)
		return
	}
	url := "ws://fancy-chat.io/ws&token=one-time-token"
	resp := &LogResponse{url}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error (json Encoder)", 500)
		return
	}
}
