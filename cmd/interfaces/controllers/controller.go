package controllers

import "github.com/gin-gonic/gin"

type Controller struct{}

// SendJson 正常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) SendJson(ctx *gin.Context, status int, data map[string]interface{}) {
	ctx.JSON(status, data)
}

// SendErrorJson 異常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) SendErrorJson(ctx *gin.Context, status int, errorMessages []string) {
	ctx.JSON(status, gin.H{"errors": errorMessages})
}
