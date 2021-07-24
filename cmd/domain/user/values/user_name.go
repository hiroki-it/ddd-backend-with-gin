package values

import "github.com/hiroki-it/ddd-api-with-go-gin/cmd/domain"

type UserName struct {
	firstName     string
	lastName      string
	firstKanaName string
	lastKanaName  string
}

var _ domain.Value = &UserName{}

// NewUserName コンストラクタ
func NewUserName(firstName string, lastName string, firstKanaName string, lastKanaName string) *UserName {
	return &UserName{
		firstName:     firstName,
		lastName:      lastName,
		firstKanaName: firstKanaName,
		lastKanaName:  lastKanaName,
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

// FullName フルネームを返却します．
func (un *UserName) FullName() string {
	return un.firstName + un.lastName
}

// FullKanaName フルネームを返却します．
func (un *UserName) FullKanaName() string {
	return un.firstKanaName + un.lastKanaName
}
