package values

type UserGenderType int

const (
	MALE UserGenderType = iota + 1
	FEMALE
)

// String 区分値を返却します．
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

// ToPrimitive ユーザ性別タイプをプリミティブ型に変換します．
func (ugt UserGenderType) ToPrimitive() int {
	return int(ugt)
}
