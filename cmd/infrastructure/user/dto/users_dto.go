package dto

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
)

// UserDTO NOTE: 利便性のため，DTOはパブリックフィールドとします．
type UserDTO struct {
	UserId            int `gorm:"primaryKey"`
	UserLastName      string
	UserFirstName     string
	UserLastKanaName  string
	UserFirstKanaName string
	UserGenderType    int
}

type UsersDTO []*UserDTO

var _ infrastructure.DTO = &UserDTO{}

// ToUser DTOをユーザエンティティに変換します．
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
		users[i] = ud.ToUser()
	}

	return users
}
