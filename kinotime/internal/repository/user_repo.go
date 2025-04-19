package repository

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(ctx context.Context, username, password string) error {
	if err := repo.DB.Ping(); err != nil {
		log.Println("Database is closed or unreachable:", err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}

	_, err = repo.DB.ExecContext(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	return err
}

func (repo *UserRepository) AuthenticateUser(ctx context.Context, username, password string) (string, bool) {
	if err := repo.DB.Ping(); err != nil {
		log.Println("Database is closed or unreachable:", err)
		return "", false
	}

	var storedPassword string
	err := repo.DB.QueryRowContext(ctx, "SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		return "", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return "", false
	}

	return storedPassword, true
}

func (repo *UserRepository) GetUserByUsername(ctx context.Context, username string) (string, bool) {
	if err := repo.DB.Ping(); err != nil {
		log.Println("Database is closed or unreachable:", err)
		return "", false
	}

	var storedPassword string
	err := repo.DB.QueryRowContext(ctx, "SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		return "", false
	}
	return storedPassword, true
}
