package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/db"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/logger"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/routers"
)

func main() {

	// データベースに接続します．
	DB, err := db.NewDB()

	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger()

	if err != nil {
		panic(err)
	}

	// 最後にデータベースとの接続を切断します．
	defer func(DB *db.DB) {
		err := DB.Close()
		if err != nil {
			log.Log().Error(err.Error())
		}
	}(DB)

	err = DB.AutoMigrate()

	if err != nil {
		log.Log().Fatal(err.Error())
	}

	// コントローラにルーティングします．
	router := routers.NewRouter(gin.Default(), DB)

	err = router.Run()

	if err != nil {
		log.Log().Error(err.Error())
	}
}
