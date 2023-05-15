package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/user", UserCreate)
	// http.HandleFunc("/user/login", UserLogin)
	// initialize our database connection
	// initDB()
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// func UserCreate(w http.ResponseWriter, r *http.Request) {
// 	os.WriteFile("t.txt", )
// }

func UserCreate(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance

	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	fmt.Println(creds)

	if err != nil {
		fmt.Println("Error")

		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	os.WriteFile("t.txt", hashedPassword, 0600)
	fmt.Fprintf(w, "for passowrd %s hash is %s", creds.Password, hashedPassword)
	// // Next, insert the username, along with the hashed password into the database
	// if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
	// 	// If there is any issue with inserting into the database, return a 500 error
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// func initDB() {
// 	var err error
// 	// Connect to the postgres db
// 	//you might have to change the connection string to add your database credentials
// 	db, err = sql.Open("postgres", "dbname=mydb sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// }
