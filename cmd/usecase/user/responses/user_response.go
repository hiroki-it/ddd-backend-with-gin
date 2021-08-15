package responses

type UserGetResponse struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserKanaName   string `json:"kana_name"`
	UserGenderType string `json:"gender_type"`
}

type UserCreateResponse struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserKanaName   string `json:"kana_name"`
	UserGenderType string `json:"gender_type"`
}

type UserUpdateResponse struct {
	UserId         int    `json:"id"`
	UserName       string `json:"name"`
	UserGenderType string `json:"gender_type"`
}

type UserDeleteResponse struct {
	UserId int `json:"id"`
}
