package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) RegisterUser(name, email string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	query := `
		INSERT INTO users(name, email)
		VALUES($1, $2)
		RETURNING id
	`

	var userID int
	err := s.db.QueryRow(query, name, email).Scan(&userID)
	if err != nil {
		return err
	}

	log.Printf("User created: %d", userID)

	emailBody := fmt.Sprintf(
		"Welcome %s! Thanks for registering.",
		name,
	)

	fmt.Printf(
		"Sending email to %s: %s\n",
		email,
		emailBody,
	)

	return nil
}

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://user:pass@localhost/app?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	service := NewUserService(db)

	err = service.RegisterUser(
		"Aditya",
		"aditya@example.com",
	)
	if err != nil {
		log.Fatal(err)
	}
}
