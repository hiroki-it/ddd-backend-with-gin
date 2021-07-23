package domain

type ValueObject interface {

	// Equals 等価性を検証します．
	Equals(target ValueObject) bool
}
