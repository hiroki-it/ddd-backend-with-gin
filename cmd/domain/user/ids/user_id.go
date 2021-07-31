package ids

type UserId int

// ToPrimitive UserIdをプリミティブ型に変換します．
func (ui UserId) ToPrimitive() int {
	return int(ui)
}
