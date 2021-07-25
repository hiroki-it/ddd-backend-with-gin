package converters

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"errors": err})
			return
		}

		ctx.Set("id", id)

		ctx.Next()
	}
}
