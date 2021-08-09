package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

// ErrorJSON 異常系レスポンスを返却します．各APIのレスポンスのデータ構造を統一します．
func (c *Controller) ErrorJSON(context *gin.Context, status int, errors []string) {
	context.JSON(status, gin.H{"errors": errors})
}
