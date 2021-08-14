package interactors

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/boundaries"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/requests"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/responses"
)

type UserInteractor struct {
	userRepository repositories.UserRepository
}

var _ boundaries.UserInputBoundary = &UserInteractor{}

// NewUserInteractor コンストラクタ
func NewUserInteractor(userRepository repositories.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

// GetUser ユーザを取得します．
func (ui *UserInteractor) GetUser(guRequest *requests.GetUserRequest) (*responses.GetUserResponse, error) {
	user, err := ui.userRepository.FindById(ids.UserId(guRequest.UserId))

	if err != nil {
		return nil, err
	}

	return &responses.GetUserResponse{
		UserId:         user.Id().ToPrimitive(),
		UserName:       user.Name().FullName(),
		UserKanaName:   user.Name().FullKanaName(),
		UserGenderType: user.GenderType().String(),
	}, nil
}

// CreateUser ユーザを作成します．
func (ui *UserInteractor) CreateUser(cuRequest *requests.CreateUserRequest) (*responses.CreateUserResponse, error) {
	user := entities.NewUser(
		ids.UserId(cuRequest.UserId),
		values.NewUserName(cuRequest.UserLastName, cuRequest.UserFirstName, cuRequest.UserLastKanaName, cuRequest.UserFirstKanaName),
		values.UserGenderType(cuRequest.UserGenderType),
	)

	err := ui.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &responses.CreateUserResponse{
		UserId:         user.Id().ToPrimitive(),
		UserName:       user.Name().FullName(),
		UserKanaName:   user.Name().FullKanaName(),
		UserGenderType: user.GenderType().String(),
	}, nil
}
