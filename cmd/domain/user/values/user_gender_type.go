package values

type UserGenderType int

const (
	MALE UserGenderType = iota + 1
	FEMALE
)

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
