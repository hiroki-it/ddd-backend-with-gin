package repositories

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
)

type UserRepository interface {
	FindById(id ids.UserId) (*entities.User, error)
	FindAll() (entities.Users, error)
	Update() (*entities.User, error)
}
