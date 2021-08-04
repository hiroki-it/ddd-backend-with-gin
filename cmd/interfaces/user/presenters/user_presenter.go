package presenters

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces"
)

type getUserPresenter struct {
	UserId int `json:"id"`
}

type createUserPresenter struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserGenderType int    `json:"gender_type"`
}

type updateUserPresenter struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserGenderType int    `json:"gender_type"`
}

type deleteUserPresenter struct {
	UserId int `json:"id"`
}

// ToGetUserPresenter 取得レスポンスデータを作成します．
func ToGetUserPresenter(user *entities.User) *getUserPresenter {
	return &getUserPresenter{
		UserId: int(user.Id()),
	}
}

// ToCreateUserPresenter 作成レスポンスデータを作成します．
func ToCreateUserPresenter(user *entities.User) *createUserPresenter {
	return &createUserPresenter{
		UserId:         int(user.Id()),
		UserName:       user.Name().FullName(),
		UserGenderType: user.GenderType().String(),
	}
}

var _, _, _, _ interfaces.Presenter = &getUserPresenter{},
	&createUserPresenter{},
	&updateUserPresenter{},
	&deleteUserPresenter{}
