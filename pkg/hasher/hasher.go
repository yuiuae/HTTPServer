// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package hasher

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hash for the password...
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hash), err
}

// CheckPasswordHash checks password by hash...
func CheckPasswordHash(password, hash string) bool {
	return false
}
