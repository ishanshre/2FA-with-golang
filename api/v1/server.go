package v1

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/ishanshre/2FA-with-golang/api/v1/models"
)

type Storage interface {
	UserSignUp(*models.RegisterUser) error
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
			otp_auth_url VARCHAR(255)
		);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) UserSignUp(user *models.RegisterUser) error {
	return nil
}
