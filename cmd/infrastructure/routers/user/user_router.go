package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"

	repositories "github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/repositories/user"
	controllers "github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/controllers/user"
	usecase "github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/usecases/user"
)

// UserRouter Userに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *infrastructure.DB) {

	userRouter := router.Group("/users")
	{
		c := controllers.NewUserController(usecase.NewUserUsecase(repositories.NewUserRepository(db)))
		userRouter.GET("/:id", c.GetUser)
	}
}
