package domain

// ID ゲッターとの名前の重複を防ぐために，名前を大文字にしています．
type ID interface {

	// Id idを返却します．
	Id() string

	// Equals 等価性を検証します．
	Equals(target ID) bool
}
