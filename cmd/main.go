package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/routers/user"
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

	route := gin.Default()

	// コントローラにルーティングします．
	user.UserRouter(route, db)

	err = route.Run()

	if err != nil {
		logger.Log().Fatal(err.Error())
	}
}
