package values

import "github.com/hiroki-it/ddd-backend-with-gin/cmd/domain"

type UserName struct {
	lastName      string
	firstName     string
	lastKanaName  string
	firstKanaName string
}

var _ domain.Value = &UserName{}

// NewUserName コンストラクタ
func NewUserName(lastName string, firstName string, lastKanaName string, firstKanaName string) *UserName {
	return &UserName{
		lastName:      lastName,
		firstName:     firstName,
		lastKanaName:  lastKanaName,
		firstKanaName: firstKanaName,
	}
}

// Equal 等価性を検証します．
func (un *UserName) Equal(target domain.Value) bool {
	t := target.(*UserName)

	// 全ての属性が一致した場合
	if un.firstName == t.firstName &&
		un.lastName == t.lastName &&
		un.firstKanaName == t.firstKanaName &&
		un.lastKanaName == t.lastKanaName {
		return true
	}

	return false
}

// LastName 苗字を返却します．
func (un *UserName) LastName() string {
	return un.lastName
}

// FirstName 名前を返却します．
func (un *UserName) FirstName() string {
	return un.firstName
}

// LastKanaName 苗字カナを返却します．
func (un *UserName) LastKanaName() string {
	return un.lastKanaName
}

// FirstKanaName 名前カナを返却します．
func (un *UserName) FirstKanaName() string {
	return un.firstKanaName
}

// FullName フルネームを返却します．
func (un *UserName) FullName() string {
	return un.firstName + un.lastName
}

// FullKanaName フルネームを返却します．
func (un *UserName) FullKanaName() string {
	return un.firstKanaName + un.lastKanaName
}
