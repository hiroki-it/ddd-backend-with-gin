package inputs

type UserInput struct {
	userId         int
	userName       string
	userGenderType int
}

// NewUserInput コンストラクタ
func NewUserInput(userId int, userName string, userGenderType int) *UserInput {

	return &UserInput{
		userId:         userId,
		userName:       userName,
		userGenderType: userGenderType,
	}
}
