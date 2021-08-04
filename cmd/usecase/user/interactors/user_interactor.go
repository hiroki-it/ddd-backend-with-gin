package interactors

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/values"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/user/presenters"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/inputs"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/outputs"
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
func (ui *UserInteractor) GetUser(input *inputs.GetUserInput) (*presenters.GetUserPresenter, error) {
	user, err := ui.userRepository.FindById(ids.UserId(input.UserId))

	if err != nil {
		return nil, err
	}

	guo := &outputs.GetUserOutput{
		UserId:         user.Id().ToPrimitive(),
		UserName:       user.Name().FullName(),
		UserKanaName:   user.Name().FullKanaName(),
		UserGenderType: user.GenderType().String(),
	}

	return presenters.ToGetUserPresenter(guo), nil
}

// CreateUser ユーザを作成します．
func (ui *UserInteractor) CreateUser(cui *inputs.CreateUserInput) (*presenters.CreateUserPresenter, error) {
	user := entities.NewUser(
		ids.UserId(cui.UserId),
		values.NewUserName(cui.UserLastName, cui.UserFirstName, cui.UserLastKanaName, cui.UserFirstKanaName),
		values.UserGenderType(cui.UserGenderType),
	)

	user, err := ui.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	cuo := &outputs.CreateUserOutput{
		UserId:         user.Id().ToPrimitive(),
		UserName:       user.Name().FullName(),
		UserKanaName:   user.Name().FullKanaName(),
		UserGenderType: user.GenderType().String(),
	}

	return presenters.ToCreateUserPresenter(cuo), nil
}
