package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	db, err := sql.Open("postgres", AppConfig.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(2 * time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatalf("PostgreSQL ping failed: %v", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL.")
}
