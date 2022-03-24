package database

import (
	"database/sql"
	"fmt"
	"log"
	"os" // used to read the environment variable

	"github.com/joho/godotenv" // package used to read the .env file
)

type DataBaseConnection interface {
	Create() *sql.DB
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
