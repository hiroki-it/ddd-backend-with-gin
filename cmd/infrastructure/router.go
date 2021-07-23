package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	route *gin.Engine
}

// NewRouter コンストラクタ
func NewRouter() *Router {

	r := gin.Default()

	return &Router{
		route: r,
	}
}

// Route routeを返却します．
func (r *Router) Route() *gin.Engine {
	return r.route
}
