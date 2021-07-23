package values

type UserGenderType int

const (
	MALE UserGenderType = iota + 1
	FEMALE
)

func NewUserGenderType(userGenderType int) UserGenderType {
	return UserGenderType(userGenderType)
}

func (ugt UserGenderType) String() string {

	switch ugt {
	case MALE:
		return "男性"
	case FEMALE:
		return "女性"
	default:
		return "未回答"
	}
}
