package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/user/controllers"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/user/presenters"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/interactors"
)

// UserRouter ユーザに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *infrastructure.DB) {
	userRouter := router.Group("/users")
	{
		c := controllers.NewUserController(interactors.NewUserInteractor(repositories.NewUserRepository(db), &presenters.UserPresenter{}))
		userRouter.GET("/user/:id", c.GetUser)
	}
}
