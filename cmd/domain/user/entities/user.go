package entities

import (
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/domain"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/domain/user/values"
)

type User struct {
	userId         ids.UserId
	userName       *values.UserName
	userGenderType values.UserGenderType
}

type Users []*User

var _ domain.Entity = &User{}

// NewUser コンストラクタ
func NewUser(userId ids.UserId, userName *values.UserName, userGenderType values.UserGenderType) *User {
	return &User{
		userId:         userId,
		userName:       userName,
		userGenderType: userGenderType,
	}
}

// Id IDを返却します．
func (u *User) Id() ids.UserId {
	return u.userId
}

// Name 名前を返却します．
func (u *User) Name() *values.UserName {
	return u.userName
}

// GenderType 性別を返却します．
func (u *User) GenderType() values.UserGenderType {
	return u.userGenderType
}

// Equal 等価性を検証します．
func (u *User) Equal(target domain.Entity) bool {
	return u.userId == target.(*User).Id()
}
