package inputs

type GetUserInput struct {
	UserId int `binding:"required,min=1"`
}

type CreateUserInput struct {
	UserId            int    // NOTE: 何も値を設定しないことにより，ゼロが割り当てられます．
	UserLastName      string `json:"last_name" binding:"required"`
	UserFirstName     string `json:"first_name" binding:"required"`
	UserLastKanaName  string `json:"last_kana_name" binding:"required"`
	UserFirstKanaName string `json:"first_kana_name" binding:"required"`
	UserGenderType    int    `json:"gender_type" binding:"required,min=1"`
}

type UpdateUserInput struct {
	UserId int `json:"id" binding:"required,min=1"`
}

type DeleteUserInput struct {
	UserId int `json:"id" binding:"required,min=1"`
}
