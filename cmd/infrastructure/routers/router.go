package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure"
	"github.com/hiroki-it/ddd-api-with-go-gin/cmd/infrastructure/middlewares"
)

type Router struct {
	router *gin.Engine
	db     *infrastructure.DB
}

// NewRouter コンストラクタ
func NewRouter(router *gin.Engine, db *infrastructure.DB) *Router {
	return &Router{
		router: router,
		db:     db,
	}
}

// Run 全てのルーティングを実行します．
func (r *Router) Run() error {
	r.router.Use(
		middlewares.HandleError(),
	)

	UserRouter(r.router, r.db)

	return r.router.Run()
}
