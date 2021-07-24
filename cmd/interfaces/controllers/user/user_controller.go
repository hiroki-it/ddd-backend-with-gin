package controllers

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"

	usecases "github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/usecases/user"
)

type UserController struct {
	userUsecase *usecases.UserUsecase
}

// NewUserController コンストラクタ
func NewUserController(userUsecase *usecases.UserUsecase) *UserController {

	return &UserController{
		userUsecase: userUsecase,
	}
}

// GetUser 単一のユーザを取得します．
func (ctl *UserController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := ctl.userUsecase.GetUser(ids.UserId(id))

	if err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	c.JSON(200, user)
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
