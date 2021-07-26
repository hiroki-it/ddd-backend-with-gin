package usecases

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"

	inputs "github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/inputs/user"
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
func (uu *UserInteractor) GetUser(input *inputs.GetUserInput) (*entities.User, error) {
	return nil, nil
}
