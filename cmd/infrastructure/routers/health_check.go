package routers

import (
	"github.com/gin-gonic/gin"
)

// HealthCheckRouter ルーティングを実行します．
func HealthCheckRouter(router *gin.Engine) {
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	return
}
