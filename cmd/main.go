package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/routers"
)

func main() {

	// データベースに接続します．
	db, err := infrastructure.NewDB()

	if err != nil {
		panic(err)
	}

	logger, err := infrastructure.NewLogger()

	if err != nil {
		panic(err)
	}

	// 最後にデータベースとの接続を切断します．
	defer func(db *infrastructure.DB) {
		err := db.Close()
		if err != nil {
			logger.Log().Error(err.Error())
		}
	}(db)

	err = db.AutoMigrate()

	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	// コントローラにルーティングします．
	router := routers.NewRouter(gin.Default(), db)

	err = router.Run()

	if err != nil {
		logger.Log().Error(err.Error())
	}
}
