package main

import (
	"log"
	"os"

	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
)

func main() {

	db, err := infrastructure.NewDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.AutoMigrate()

	if err != nil {
		log.Fatal(err)
	}
}
