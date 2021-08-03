package interactor

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/input"
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
func (ui *UserInteractor) GetUser(input *input.GetUserInput) (*entities.User, error) {
	return ui.userRepository.FindById(ids.UserId(input.UserId))
}
