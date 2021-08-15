package boundaries

import (
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/requests"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/usecase/user/responses"
)

type UserInputBoundary interface {
	GetUser(*requests.UserGetRequest) (*responses.UserGetResponse, error)
	CreateUser(*requests.UserCreateRequest) (*responses.UserCreateResponse, error)
}
