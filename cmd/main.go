package main

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
)

func main() {

	logger := infrastructure.NewLogger()

	// データベースに接続します．
	db, err := infrastructure.NewDB()

	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	defer db.Close()

	err = db.AutoMigrate()

	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	// コントローラにルーティングします．
	err = infrastructure.NewRouter().Route().Run()

	if err != nil {
		logger.Log().Fatal(err.Error())
	}
}
