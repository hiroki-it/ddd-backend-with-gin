package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

// NewRouter コンストラクタ
func NewRouter() *Router {

	r := gin.Default()

	return &Router{
		engine: r,
	}
}

// Engine engineを返却します．
func (r *Router) Engine() *gin.Engine {
	return r.engine
}
