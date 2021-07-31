package repository

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/user/dto"
)

type UserRepository struct {
	db *infrastructure.DB
}

var _ repositories.UserRepository = &UserRepository{}

// NewUserRepository コンストラクタ
func NewUserRepository(db *infrastructure.DB) repositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// FindById IDを元にユーザを返却します．
func (ur *UserRepository) FindById(id ids.UserId) (*entities.User, error) {
	userDTO := &dto.UserDTO{}

	err := ur.db.Conn().First(userDTO, id.ToPrimitive()).Error

	if err != nil {
		return nil, err
	}

	return userDTO.ToUser(), nil
}

// FindAll 全てのユーザを返却します．
func (ur *UserRepository) FindAll() (entities.Users, error) {
	usersDTO := &dto.UsersDTO{}

	err := ur.db.Conn().First(usersDTO).Error

	if err != nil {
		return nil, err
	}

	return usersDTO.ToUsers(), nil
}
