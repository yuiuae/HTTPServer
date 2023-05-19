// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hash for the password...
func HashPassword(password string) (string, error) {
	// for i := 0; i < 10; i++ {
	// 	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	// 	fmt.Printf("%s\n", hash)
	// }

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hash), err
}

// CheckPasswordHash checks password by hash...
func CheckPasswordHash(password, hash string) bool {
	fmt.Println(password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
	// if err != nil {
	// 	return false
	// }
	// return true
}
