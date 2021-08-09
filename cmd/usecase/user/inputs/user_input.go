package inputs

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/requests"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/responses"
)

type UserInput interface {
	GetUser(*requests.GetUserRequest) (*responses.GetUserResponse, error)
	CreateUser(*requests.CreateUserRequest) (*responses.CreateUserResponse, error)
}
