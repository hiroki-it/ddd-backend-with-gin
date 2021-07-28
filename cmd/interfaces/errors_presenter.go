package interfaces

// errorsPresenter 異常系レスポンスを定義します．
type errorsPresenter struct {
	Errors []string `json:"errors"`
}

// ToErrorsPresenter 変換します．
func ToErrorsPresenter(errors []string) *errorsPresenter {
	return &errorsPresenter{
		Errors: errors,
	}
}
