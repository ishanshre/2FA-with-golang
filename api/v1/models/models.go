package models

import (
	"database/sql"
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
	ID           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Password     string         `json:"-"`
	Otp_enabled  bool           `json:"otp_enabled"`
	Otp_verified bool           `json:"otp_verified"`
	Otp_secret   sql.NullString `json:"otp_secret"`
	Otp_auth_url sql.NullString `json:"otp_auth_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	LastLogin    time.Time      `json:"last_login"`
}

// for Login Request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Password struct {
	Password string `json:"password"`
}

type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type ScanSecret struct {
	Secret string `json:"secret"`
}

type GetSecret struct {
	Username   string `json:"username"`
	Otp_secret string `json:"otp_secret"`
}

type Username struct {
	UserName string `json:"username"`
}

type GenerateOTP struct {
	Otp_secret   string `json:"otp_secret"`
	Otp_auth_url string `json:"otp_auth_url"`
}

type Update struct {
	Username     string `json:"username"`
	Otp_secret   string `json:"otp_secret"`
	Otp_auth_url string `json:"otp_auth_url"`
}

type OTP struct {
	Token      string `json:"token"`
	Otp_secret string `json:"otp_secret"`
}

//

func NewRegisterUser(name, username, email, password string) *RegisterUser {
	id := uuid.New()
	now := time.Now()
	zero := time.Time{}
	return &RegisterUser{
		ID:        id,
		Name:      name,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
		LastLogin: zero,
	}
}
