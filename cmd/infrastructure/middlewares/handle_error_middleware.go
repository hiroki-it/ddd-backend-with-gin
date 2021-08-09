package middlewares

import "github.com/gin-gonic/gin"

// HandleError ミドルウェア処理中のエラーをハンドリングします．
func HandleError() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		// 処理中の全てのエラーを取得します．
		errs := context.Errors

		// 処理中にエラーが発生していた場合
		if errs != nil {
			context.JSON(400, gin.H{"errors": errs.Errors()})
			return
		}

		// エラーが発生していなかった場合
		return
	}
}
