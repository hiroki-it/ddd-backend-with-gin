package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/user/repository"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/user/controller"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/interactor"
)

// UserRouter ユーザに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *infrastructure.DB) {
	userRouter := router.Group("/users")
	{
		c := controller.NewUserController(interactor.NewUserInteractor(repository.NewUserRepository(db)))
		userRouter.GET("/:id", c.GetUser)
	}
}
