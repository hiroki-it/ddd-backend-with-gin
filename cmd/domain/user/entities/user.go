package entities

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
)

type User struct {
	id         ids.UserId
	name       values.UserName
	genderType values.UserGenderType
}
