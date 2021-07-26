package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"

	repositories "github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/repositories/user"
	controllers "github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/controllers/user"
	interactors "github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/interactors/user"
)

// UserRouter ユーザに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *infrastructure.DB) {
	userRouter := router.Group("/users")
	{
		c := controllers.NewUserController(interactors.NewUserInteractor(repositories.NewUserRepository(db)))
		userRouter.GET("/:id", c.GetUser)
	}
}
