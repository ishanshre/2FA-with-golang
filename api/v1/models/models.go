package models

import (
	"time"

	"github.com/google/uuid"
)

// For creating new table in the database
type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Otp_enabled  bool      `json:"otp_enabled"`
	Otp_verified bool      `json:"otp_verified"`
	Otp_secret   string    `json:"otp_secret"`
	Otp_auth_url string    `json:"otp_auth_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	LastLogin    time.Time `json:"last_login"`
}

// For Registering new user
type RegisterUser struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LastLogin time.Time `json:"last_login"`
}

// for get endpoint
type UserProfile struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Otp_enabled  bool      `json:"otp_enabled"`
	Otp_verified bool      `json:"otp_verified"`
	Otp_secret   string    `json:"otp_secret"`
	Otp_auth_url string    `json:"otp_auth_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	LastLogin    time.Time `json:"last_login"`
}

// for Login Request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//

func NewRegisterUser(name, username, email, password string) (*RegisterUser, error) {
	id := uuid.New()
	return &RegisterUser{
		ID:       id,
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}
