package v1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishanshre/2FA-with-golang/api/v1/models"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	/*
		Response writer middleware
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandler(f ApiFunc) http.HandlerFunc {
	/*
		Middleware that returns HandleFunc
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func checkPostMethod(r *http.Request) error {
	if r.Method == "POST" {
		return nil
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}

func scanPassword(rows *sql.Rows) (*models.Password, error) {
	password := new(models.Password)
	err := rows.Scan(
		&password.Password,
	)
	return password, err
}

func scanUserProfile(rows *sql.Rows) (*models.UserProfile, error) {
	user := new(models.UserProfile)
	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Otp_enabled,
		&user.Otp_verified,
		&user.Otp_secret,
		&user.Otp_auth_url,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	)
	return user, err
}
