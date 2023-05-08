package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	db := initDB()
	migrate(db)
}

func initDB() (db *sql.DB) {

	path, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", path))
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}

func migrate(db *sql.DB) {
}
