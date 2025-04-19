package repository

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectPostgre(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Database connection established.")
	return db
}
