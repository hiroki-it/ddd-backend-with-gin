package usecases

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
)

type UserInteractor struct {
	userRepository repositories.UserRepository
}

// NewUserInteractor コンストラクタ
func NewUserInteractor(userRepository repositories.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

// GetUser ユーザを取得します．
func (uu *UserInteractor) GetUser(userId ids.UserId) (*entities.User, error) {
	return nil, nil
}
