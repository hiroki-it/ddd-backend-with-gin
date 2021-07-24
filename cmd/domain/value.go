package domain

type Value interface {
	// Equal 等価性を検証します．
	Equal(target Value) bool
}
