package user

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
)

type UserRepository struct {
	db *infrastructure.DB
}

// NewUserRepository コンストラクタ
func NewUserRepository(db *infrastructure.DB) repositories.UserRepository {

	return &UserRepository{
		db: db,
	}
}
