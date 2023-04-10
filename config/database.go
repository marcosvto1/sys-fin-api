package config

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	// CREATE TABLE USERS
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, profile INTEGER")
	if err != nil {
	}

	// CREATE TABLE ESPECIALIDADES
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS specialties (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255) NOT NULL, description TEXT")
	if err != nil {
		log.Fatal(err)
	}
}
