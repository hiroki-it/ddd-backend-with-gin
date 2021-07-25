package converters

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ConvertId パスパラメータのidのデータ型を変換します．
func ConvertId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Set("id", id)

		ctx.Next()
	}
}
