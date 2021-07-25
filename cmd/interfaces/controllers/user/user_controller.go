package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/controllers"

	usecases "github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/usecases/user"
)

type UserController struct {
	*controllers.Controller
	userUsecase *usecases.UserUsecase
}

// NewUserController コンストラクタ
func NewUserController(userUsecase *usecases.UserUsecase) *UserController {

	return &UserController{
		Controller:  &controllers.Controller{},
		userUsecase: userUsecase,
	}
}

// GetUser 単一のユーザを取得します．
func (uc *UserController) GetUser(ctx *gin.Context) {
	userId, ok := ctx.Get("id")

	if !ok {
		uc.SendErrorJson(ctx, 400, []string{"Parameters are not found."})
		return
	}

	user, err := uc.userUsecase.GetUser(userId.(ids.UserId))

	if err != nil {
		uc.SendErrorJson(ctx, 400, []string{err.Error()})
		return
	}

	uc.SendJson(ctx, 200, user)
	return
}

// GetUsers 複数のユーザを取得します．
func (ctl *UserController) GetUsers(c *gin.Context) {
}

// CreateUsers ユーザを作成します．
func (ctl *UserController) CreateUsers(c *gin.Context) {
}

// UpdateUser ユーザを更新します．
func (ctl *UserController) UpdateUser(c *gin.Context) {
}

// DeleteUser ユーザを削除します．
func (ctl *UserController) DeleteUser(c *gin.Context) {
}
