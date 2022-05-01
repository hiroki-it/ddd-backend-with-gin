package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/database"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/user/repositories"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/interfaces/user/controllers"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/usecase/user/interactors"
)

// UserRouter ユーザに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *database.DB) {
	userRouter := router.Group("/users")
	{
		c := controllers.NewUserController(interactors.NewUserInteractor(repositories.NewUserRepository(db)))
		userRouter.GET("/:id", c.GetUser)
		userRouter.POST("/", c.CreateUser)
	}
}
