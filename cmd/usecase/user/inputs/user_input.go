package inputs

type GetUserInput struct {
	UserId int `binding:"required,min=1"`
}

type CreateUserInput struct {
	UserName       string `json:"name" binding:"required"`
	UserGenderType int    `json:"gender_type" binding:"required,min=1"`
}

type UpdateUserInput struct {
	UserId int `json:"id" binding:"required,min=1"`
}

type DeleteUserInput struct {
	UserId int `json:"id" binding:"required,min=1"`
}
