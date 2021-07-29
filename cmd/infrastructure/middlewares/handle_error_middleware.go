package middlewares

import "github.com/gin-gonic/gin"

// HandleError ミドルウェア処理中のエラーをハンドリングします．
func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// 処理中の全てのエラーを取得します．
		errs := ctx.Errors

		// 処理中にエラーが発生していた場合
		if errs != nil {
			ctx.JSON(400, gin.H{"errors": errs.Errors()})
			return
		}

		// エラーが発生していなかった場合
		return
	}
}
