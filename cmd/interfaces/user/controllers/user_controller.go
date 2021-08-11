package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/boundaries"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/requests"
)

type UserController struct {
	*interfaces.Controller
	userInputBoundary boundaries.UserInputBoundary
}

// NewUserController コンストラクタ
func NewUserController(userInputBoundary boundaries.UserInputBoundary) *UserController {

	return &UserController{
		Controller:        &interfaces.Controller{},
		userInputBoundary: userInputBoundary,
	}
}

// GetUser 単一のユーザを取得します．
func (uc *UserController) GetUser(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		uc.ErrorJSON(context, 400, []string{"Parameters are not found."})
		return
	}

	guRequest := &requests.GetUserRequest{UserId: userId}

	guResponse, err := uc.userInputBoundary.GetUser(guRequest)

	if err != nil {
		uc.ErrorJSON(context, 400, []string{err.Error()})
		return
	}

	context.JSON(200, guResponse)
	return
}

// GetUsers 複数のユーザを取得します．
func (uc *UserController) GetUsers(context *gin.Context) {
}

// CreateUser ユーザを作成します．
func (uc *UserController) CreateUser(context *gin.Context) {
	cuRequest := &requests.CreateUserRequest{}

	err := context.Bind(cuRequest)

	if err != nil {
		uc.ErrorJSON(context, 400, []string{"Parameters are not found."})
		return
	}

	cuResponse, err := uc.userInputBoundary.CreateUser(cuRequest)

	if err != nil {
		uc.ErrorJSON(context, 400, []string{err.Error()})
		return
	}

	context.JSON(200, cuResponse)
	return
}

// UpdateUser ユーザを更新します．
func (uc *UserController) UpdateUser(context *gin.Context) {
}

// DeleteUser ユーザを削除します．
func (uc *UserController) DeleteUser(context *gin.Context) {
}
