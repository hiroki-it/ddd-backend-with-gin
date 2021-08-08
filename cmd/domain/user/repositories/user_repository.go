package repositories

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
)

type UserRepository interface {
	Create(*entities.User) (*entities.User, error)
	FindById(ids.UserId) (*entities.User, error)
	FindAll() (entities.Users, error)
	Update(*entities.User) (*entities.User, error)
	Delete(ids.UserId) error
}
