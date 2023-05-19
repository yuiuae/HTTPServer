package handlers

import (
	"fmt"
	"net/http"
)

func UserAll(w http.ResponseWriter, r *http.Request) {
	for key, val := range usersTable {
		fmt.Fprintf(w, "\n%s = %v, %v", key, val.Id, val.Passhash)
	}
}
