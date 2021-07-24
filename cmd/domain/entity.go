package domain

type Entity interface {
	// Equal 等価性を検証します．
	Equal(target Entity) bool
}
