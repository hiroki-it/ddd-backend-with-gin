package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/database"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/logger"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/routers"
)

func main() {

	// データベースに接続します．
	db, err := database.NewDB()

	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger()

	if err != nil {
		panic(err)
	}

	// 最後にデータベースとの接続を切断します．
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Log().Error(err.Error())
		}
	}(db)

	err = db.AutoMigrate()

	if err != nil {
		log.Log().Fatal(err.Error())
	}

	// コントローラにルーティングします．
	router := routers.NewRouter(gin.Default(), db)

	err = router.Run()

	if err != nil {
		log.Log().Error(err.Error())
	}
}
