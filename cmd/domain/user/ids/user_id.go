package ids

type UserId int

// ToPrimitive ユーザIDをプリミティブ型に変換します．
func (ui UserId) ToPrimitive() int {
	return int(ui)
}
