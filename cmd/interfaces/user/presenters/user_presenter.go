package presenters

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/outputs"
)

type UserPresenter struct {
}

type GetUserPresenter struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserKanaName   string `json:"kana_name"`
	UserGenderType string `json:"gender_type"`
}

type CreateUserPresenter struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserKanaName   string `json:"kana_name"`
	UserGenderType string `json:"gender_type"`
}

type UpdateUserPresenter struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserGenderType string `json:"gender_type"`
}

type DeleteUserPresenter struct {
	UserId int `json:"id"`
}

// ToGetUserPresenter 取得レスポンスデータを作成します．
func (up *UserPresenter) ToGetUserPresenter(guo *outputs.GetUserOutput) *GetUserPresenter {
	return &GetUserPresenter{
		UserId:         guo.UserId,
		UserName:       guo.UserName,
		UserKanaName:   guo.UserKanaName,
		UserGenderType: guo.UserGenderType,
	}
}

// ToCreateUserPresenter 作成レスポンスデータを作成します．
func (up *UserPresenter) ToCreateUserPresenter(cuo *outputs.CreateUserOutput) *CreateUserPresenter {
	return &CreateUserPresenter{
		UserId:         cuo.UserId,
		UserName:       cuo.UserName,
		UserKanaName:   cuo.UserKanaName,
		UserGenderType: cuo.UserGenderType,
	}
}

var _, _, _, _, _ interfaces.Presenter = &UserPresenter{},
	&GetUserPresenter{},
	&CreateUserPresenter{},
	&UpdateUserPresenter{},
	&DeleteUserPresenter{}
