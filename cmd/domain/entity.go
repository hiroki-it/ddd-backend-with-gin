package domain

type Entity interface {

	// Equals 等価性を検証します．
	Equals(target Entity) bool
}
