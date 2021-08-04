package outputs

type GetUserOutput struct {
	UserId         int
	UserName       string
	UserKanaName   string
	UserGenderType string
}

type CreateUserOutput struct {
	UserId         int
	UserName       string
	UserKanaName   string
	UserGenderType string
}
