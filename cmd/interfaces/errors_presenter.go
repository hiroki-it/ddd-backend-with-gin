package interfaces

// errorsPresenter 異常系レスポンスを定義します．
type errorsPresenter struct {
	Errors []string `json:"errors"`
}

// ErrorsPresenter 変換します．
func ErrorsPresenter(errors []string) *errorsPresenter {
	return &errorsPresenter{
		Errors: errors,
	}
}
