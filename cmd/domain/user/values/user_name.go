package values

type UserName string

func NewUserName(userName string) UserName {
	return UserName(userName)
}
