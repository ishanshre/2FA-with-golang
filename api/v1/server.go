package v1

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/ishanshre/2FA-with-golang/api/v1/models"
)

type Storage interface {
	UserSignUp(*models.RegisterUser) error
	UserLogin(username string) (*models.Password, error)
	UserUpdateLogin(username string) (*models.UserProfile, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgresStore, error) {
	/*
		Initializing the postgres database
	*/
	if err := godotenv.Load("./.env"); err != nil {
		return nil, fmt.Errorf("error in loading environment variables")
	}

	// connecting to database using postgres connection string from env file
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		return nil, fmt.Errorf("error in connecting to database: %s", err)
	}
	// now pinging if databsae is ready to accept connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// return the postgresql
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createTable()
}

func (s *PostgresStore) createTable() error {
	// creating user table if not exists in database
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255),
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255),
			password VARCHAR(500),
			otp_enabled BOOLEAN DEFAULT 'f',
			otp_verified BOOLEAN DEFAULT 'f',
 			otp_secret VARCHAR(255),
			otp_auth_url VARCHAR(255),
			created_at TIMESTAMPTZ,
			updated_at TIMESTAMPTZ,
			last_login TIMESTAMPTZ
		);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) UserSignUp(user *models.RegisterUser) error {
	query := `
		INSERT INTO users (
			id, name, username, email, password, created_at, updated_at, last_login
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		);
	`
	_, err := s.db.Query(
		query,
		user.ID,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.LastLogin,
	)
	if err != nil {
		return fmt.Errorf("error in creating new user: %s", err)
	}
	return nil
}

func (s *PostgresStore) UserLogin(username string) (*models.Password, error) {
	query := `SELECT password FROM users WHERE username = $1`
	rows, err := s.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPassword(rows)
	}
	return nil, nil
}

func (s *PostgresStore) UserUpdateLogin(username string) (*models.UserProfile, error) {
	update := `
		UPDATE users
		SET last_login = $2
		WHERE username = $1
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(update, username, time.Now())
	if err != nil {
		return nil, err
	}
	query := `
		SELECT * FROM users
		WHERE username = $1
	`
	rows, err := s.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanUserProfile(rows)
	}
	return nil, fmt.Errorf("username: %s does not exists", username)
}
