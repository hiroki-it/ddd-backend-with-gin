package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/routers"
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
	router := routers.NewRouter(gin.Default(), db)

	err = router.Run()

	if err != nil {
		logger.Log().Fatal(err.Error())
	}
}
