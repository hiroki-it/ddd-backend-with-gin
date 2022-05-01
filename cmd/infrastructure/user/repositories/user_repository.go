package repositories

import (
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/domain/user/entities"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/domain/user/ids"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/domain/user/repositories"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/database"
	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/user/dtos"
)

type UserRepository struct {
	db *database.DB
}

var _ repositories.UserRepository = &UserRepository{}

// NewUserRepository コンストラクタ
func NewUserRepository(db *database.DB) repositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create ユーザを作成します．
func (ur *UserRepository) Create(user *entities.User) error {
	userDTO := dtos.UserDTO{
		UserId:            user.Id().ToPrimitive(),
		UserLastName:      user.Name().LastName(),
		UserFirstName:     user.Name().FirstName(),
		UserLastKanaName:  user.Name().LastKanaName(),
		UserFirstKanaName: user.Name().FirstKanaName(),
		UserGenderType:    user.GenderType().ToPrimitive(),
	}

	err := ur.db.Create(&userDTO)

	if err != nil {
		return err
	}

	// DTOをユーザエンティティに変換します．
	return nil
}

// FindById IDを元にユーザを返却します．
func (ur *UserRepository) FindById(id ids.UserId) (*entities.User, error) {
	userDTO := dtos.UserDTO{}

	err := ur.db.Find(&userDTO, id.ToPrimitive())

	if err != nil {
		return nil, err
	}

	// DTOをユーザエンティティに変換します．
	return userDTO.ToUser(), nil
}

// FindAll 全てのユーザを返却します．
func (ur *UserRepository) FindAll() (entities.Users, error) {
	usersDTO := dtos.UsersDTO{}

	err := ur.db.FindAll(&usersDTO)

	if err != nil {
		return nil, err
	}

	// DTO配列をユーザエンティティ配列に変換します．
	return usersDTO.ToUsers(), nil
}

// Update ユーザを更新します．
func (ur *UserRepository) Update(user *entities.User) error {
	// ユーザエンティティをDTOに変換します．
	userDTO := dtos.UserDTO{
		UserId:            user.Id().ToPrimitive(),
		UserLastName:      user.Name().LastName(),
		UserFirstName:     user.Name().FirstName(),
		UserLastKanaName:  user.Name().LastKanaName(),
		UserFirstKanaName: user.Name().FirstKanaName(),
		UserGenderType:    user.GenderType().ToPrimitive(),
	}

	err := ur.db.Updates(&userDTO)

	if err != nil {
		return err
	}

	// DTOをユーザエンティティに変換します．
	return nil
}

// Delete ユーザを削除します．
func (ur *UserRepository) Delete(id ids.UserId) error {
	// ユーザエンティティをDTOに変換します．
	userDTO := dtos.UserDTO{
		UserId: id.ToPrimitive(),
	}

	err := ur.db.Delete(&userDTO)

	if err != nil {
		return err
	}

	return nil
}
