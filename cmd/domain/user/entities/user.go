package entities

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
)

type User struct {
	id         ids.UserId
	name       values.UserName
	genderType values.UserGenderType
}

type Users []*User

var _ domain.Entity = &User{}

// NewUser コンストラクタ
func NewUser(userId ids.UserId, userName values.UserName, userGenderType values.UserGenderType) *User {
	return &User{
		id:         userId,
		name:       userName,
		genderType: userGenderType,
	}
}

// Id IDを返却します．
func (u *User) Id() ids.UserId {
	return u.id
}

// Name 名前を返却します．
func (u *User) Name() values.UserName {
	return u.name
}

// GenderType 性別を返却します．
func (u *User) GenderType() values.UserGenderType {
	return u.genderType
}

// Equal 等価性を検証します．
func (u *User) Equal(target domain.Entity) bool {
	return u.id == target.(*User).Id()
}
