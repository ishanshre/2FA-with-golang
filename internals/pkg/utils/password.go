package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	/*
		This method hashes the plain string password and retirns the hash
	*/
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("error in hashing the password: %s", err)
	}
	return string(hash), nil
}

func VerifyPassword(hashedPassword, requestPassword string) error {
	/*
		Compare the hashed password from database to plain password provided by the user
	*/
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword))
}
