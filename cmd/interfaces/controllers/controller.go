package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/interfaces/presenters"
)

type Controller struct{}

// SendJson 正常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) SendJson(ctx *gin.Context, status int, presenter presenters.Presenter) {
	ctx.JSON(status, presenter)
}

// SendErrorJson 異常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) SendErrorJson(ctx *gin.Context, status int, errors []string) {
	ctx.JSON(status, presenters.ToErrorsPresenter(errors))
}
