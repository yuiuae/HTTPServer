package main

import (
	"log"
	"net/http"

	"github.com/yuiuae/HTTPServer/internal/handlers"
)

// func init() {
// 	var allusers = make(map[string]users.UserInfo)
// }

func main() {
	// "UserCreate" and "CheckUser" are handler that we will implement
	http.HandleFunc("/user", handlers.UserCreate)
	http.HandleFunc("/user/login", handlers.UserLogin)
	http.HandleFunc("/admin", handlers.UserAll)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))

}