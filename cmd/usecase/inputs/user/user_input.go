package inputs

import "github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain/user/ids"

type GetUserInput struct {
	userId ids.UserId
}

// NewGetUserInput コンストラクタ
func NewGetUserInput(userId int) *GetUserInput {
	return &GetUserInput{
		userId: ids.UserId(userId),
	}
}
