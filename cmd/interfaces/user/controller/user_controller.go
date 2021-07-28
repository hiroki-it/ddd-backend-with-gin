package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/user/presenter"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/input"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/interactor"
)

type UserController struct {
	*interfaces.Controller
	userInteractor *interactor.UserInteractor
}

// NewUserController コンストラクタ
func NewUserController(userInteractor *interactor.UserInteractor) *UserController {

	return &UserController{
		Controller:     &interfaces.Controller{},
		userInteractor: userInteractor,
	}
}

// GetUser 単一のユーザを取得します．
func (uc *UserController) GetUser(ctx *gin.Context) {
	userId, ok := ctx.Get("id")

	if !ok {
		uc.SendErrorJson(ctx, 400, []string{"Parameters are not found."})
		return
	}

	gui := &input.GetUserInput{UserId: userId.(int)}

	user, err := uc.userInteractor.GetUser(gui)

	if err != nil {
		uc.SendErrorJson(ctx, 400, []string{err.Error()})
		return
	}

	uc.SendJson(ctx, 200, presenter.ToGetUserPresenter(user))
	return
}

// GetUsers 複数のユーザを取得します．
func (uc *UserController) GetUsers(ctx *gin.Context) {
}

// CreateUsers ユーザを作成します．
func (uc *UserController) CreateUsers(ctx *gin.Context) {
}

// UpdateUser ユーザを更新します．
func (uc *UserController) UpdateUser(ctx *gin.Context) {
}

// DeleteUser ユーザを削除します．
func (uc *UserController) DeleteUser(ctx *gin.Context) {
}
