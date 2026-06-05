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

type UserRepository interface {
    CreateUser(name, email string) (userID int, err error)
}

type UserValidator struct {}
type Logger struct {}
type EmailSender struct {}

type PostgresUserRepository struct {
    db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
    return &PostgresUserRepository{
        db: db,
    }
}

func (p *PostgresUserRepository) CreateUser(name, email string) (userID int, err error) {
    query := `
		INSERT INTO users(name, email)
		VALUES($1, $2)
		RETURNING id
	`
	err := p.db.QueryRow(query, name, email).Scan(&userID)
	if err != nil {
		return
	}
	
	return
}

func (u *UserValidator) ValidateUserField(name, email string) error {
    if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	
	return nil
}

func (l *Logger) Log(text string) {
    log.Print(text)
}

func (e *EmailSender) Send(email, emailBody string) {
    fmt.Printf(
		"Sending email to %s: %s\n",
		email,
		emailBody,
	)
}

type UserService struct {
	repo UserRepository
	validator *UserValidator
	logger *Logger
    emailSender *EmailSender
}

func NewUserService(
    repo UserRepository, 
    validator *UserValidator,
    logger *Logger,
    emailSender *EmailSender,
) *UserService {
	return &UserService{
	    repo: repo,
	    validator: validator,
	    logger: logger,
	    emailSender: emailSender,
	}
}

func (s *UserService) RegisterUser(name, email string) error {
	err := s.validator.ValidateUserField(name, email)
    if err != nil {
	    return err
	}
	
	userID, err := s.repo.CreateUser(name, email)
	if err != nil {
	    return err
	}
    
    s.logger.Log(fmt.Sprintf("User created: %d", userID))

	emailBody := fmt.Sprintf("Welcome %s! Thanks for registering.", name)
	s.emailSender.Send(email, emailBody)

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
	
	repo := NewPostgresUserRepository(db)
	validator := &UserValidator{}
	logger := &Logger{}
	emailSender := &EmailSender{}
	
	service := NewUserService(repo, validator, logger, emailSender)

	err = service.RegisterUser(
		"Aditya",
		"aditya@example.com",
	)
	if err != nil {
		log.Fatal(err)
	}
}
