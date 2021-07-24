package validators

type UserValidator struct {
	UserId         int    `json:"id" binding:"required,min=1"`
	UserName       string `json:"name" binding:"required"`
	UserGenderType int    `json:"gender_type" binding:"required,min=1"`
}
