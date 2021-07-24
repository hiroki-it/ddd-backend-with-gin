package values

type UserName struct {
	firstName     string
	lastName      string
	firstKanaName string
	lastKanaName  string
}

// NewUserName コンストラクタ
func NewUserName(firstName string, lastName string, firstKanaName string, lastKanaName string) *UserName {

	return &UserName{
		firstName:     firstName,
		lastName:      lastName,
		firstKanaName: firstKanaName,
		lastKanaName:  lastKanaName,
	}
}

// FullName フルネームを返却します．
func (un *UserName) FullName() string {
	return un.firstName + un.lastName
}

// FullKanaName フルネームを返却します．
func (un *UserName) FullKanaName() string {
	return un.firstKanaName + un.lastKanaName
}
