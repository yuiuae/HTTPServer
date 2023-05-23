package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yuiuae/HTTPServer/internal/handlers"
)

func main() {
	os.Setenv("PORT", "8080")
	// "UserCreate" and "CheckUser" are handler that we will implement
	http.HandleFunc("/user", handlers.UserCreate)
	http.HandleFunc("/user/login", handlers.UserLogin)
	http.HandleFunc("/admin", handlers.GetUserAll)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8080", nil))

}
