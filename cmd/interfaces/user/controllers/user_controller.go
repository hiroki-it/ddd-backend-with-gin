package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/user/presenters"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/inputs"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/interactors"
	"strconv"
)

type UserController struct {
	*interfaces.Controller
	userInteractor *interactors.UserInteractor
}

// NewUserController コンストラクタ
func NewUserController(userInteractor *interactors.UserInteractor) *UserController {

	return &UserController{
		Controller:     &interfaces.Controller{},
		userInteractor: userInteractor,
	}
}

// GetUser 単一のユーザを取得します．
func (uc *UserController) GetUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		uc.SendErrorJson(ctx, 400, []string{"Parameters are not found."})
		return
	}

	gui := &inputs.GetUserInput{UserId: userId}

	user, err := uc.userInteractor.GetUser(gui)

	if err != nil {
		uc.SendErrorJson(ctx, 400, []string{err.Error()})
		return
	}

	uc.SendJson(ctx, 200, presenters.ToGetUserPresenter(user))
	return
}

// GetUsers 複数のユーザを取得します．
func (uc *UserController) GetUsers(ctx *gin.Context) {
}

// CreateUser ユーザを作成します．
func (uc *UserController) CreateUser(ctx *gin.Context) {
	cui := &inputs.CreateUserInput{}
	err := ctx.Bind(cui)

	if err != nil {
		uc.SendErrorJson(ctx, 400, []string{"Parameters are not found."})
		return
	}

	user, err := uc.userInteractor.CreateUser(cui)

	if err != nil {
		uc.SendErrorJson(ctx, 400, []string{err.Error()})
		return
	}

	uc.SendJson(ctx, 200, presenters.ToCreateUserPresenter(user))
	return
}

// UpdateUser ユーザを更新します．
func (uc *UserController) UpdateUser(ctx *gin.Context) {
}

// DeleteUser ユーザを削除します．
func (uc *UserController) DeleteUser(ctx *gin.Context) {
}
