package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

// SendJson 正常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) SendJson(ctx *gin.Context, status int, presenter Presenter) {
	ctx.JSON(status, presenter)
}

// SendErrorJson 異常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) SendErrorJson(ctx *gin.Context, status int, errors []string) {
	ctx.JSON(status, ErrorsPresenter(errors))
}
