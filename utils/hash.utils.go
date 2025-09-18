package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// hashHelper is an unexported struct holding password helper methods.
type hashHelper struct{}

// Hash is the exported instance of hashHelper.
// You can use it like: utils.Hash.Generate("password123")
var Hash = hashHelper{}

// Generate creates a bcrypt hash of the given password.
// bcrypt.DefaultCost = 10 by default, can be increased for more security.
func (hashHelper) Generate(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Compare checks if the plain-text password matches the hashed one.
// Returns true if they match, false otherwise.
func (hashHelper) Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
