package db

import (
	"database/sql"
	"log"

	"github.com/bytebeatz/bandroom-cms/config"
)

func GetDB() *sql.DB {
	return config.DB
}

func CloseDB() {
	if config.DB != nil {
		err := config.DB.Close()
		if err != nil {
			log.Printf("Error closing DB: %v", err)
		} else {
			log.Println("PostgreSQL connection closed.")
		}
	}
}
