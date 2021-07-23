package usecases

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
)

type UserUsecase struct {
	userRepository repositories.UserRepository
}

// NewUserUsecase コンストラクタ
func NewUserUsecase(userRepository repositories.UserRepository) *UserUsecase {

	return &UserUsecase{
		userRepository: userRepository,
	}
}

// GetUser ユーザを取得します．
func (uu *UserUsecase) GetUser(userId ids.UserId) (*entities.User, error) {
	return nil, nil
}
