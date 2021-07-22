package main

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
)

func main() {

	log, err := infrastructure.NewLogger()

	if err != nil {
		panic(err)
	}

	// データベースに接続します．
	db, err := infrastructure.NewDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}
}
