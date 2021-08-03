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

// ToGetUserPresenter 変換します．
func ToGetUserPresenter(user *entities.User) *getUserPresenter {
	return &getUserPresenter{
		UserId: int(user.Id()),
	}
}

var _, _, _, _ interfaces.Presenter = &getUserPresenter{},
	&createUserPresenter{},
	&updateUserPresenter{},
	&deleteUserPresenter{}
