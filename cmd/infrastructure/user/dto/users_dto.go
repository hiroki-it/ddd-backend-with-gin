package dto

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
)

type UserDTO struct {
	Id         ids.UserId `gorm:"primaryKey"`
	Name       values.UserName
	GenderType values.UserGenderType
}

type UsersDTO []*UserDTO

// ToUser UserDTOをUserに変換します．
func (ud UserDTO) ToUser() *entities.User {

	user := entities.NewUser(
		ud.Id,
		ud.Name,
		ud.GenderType,
	)

	return user
}

// ToUsers UsersDTOをUsersに変換します．
func (usd UsersDTO) ToUsers() entities.Users {

	users := entities.Users{}

	for i, ud := range usd {
		users[i] = entities.NewUser(
			ud.Id,
			ud.Name,
			ud.GenderType,
		)
	}

	return users
}
