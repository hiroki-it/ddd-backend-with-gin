package responses

type GetUserResponse struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserKanaName   string `json:"kana_name"`
	UserGenderType string `json:"gender_type"`
}

type CreateUserResponse struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserKanaName   string `json:"kana_name"`
	UserGenderType string `json:"gender_type"`
}

type UpdateUserResponse struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserGenderType string `json:"gender_type"`
}

type DeleteUserResponse struct {
	UserId int `json:"id"`
}
