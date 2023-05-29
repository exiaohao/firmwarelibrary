package router

import (
	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	app := r.Group("/app")
	app.GET("/healthz", controller.Healthz)

	query := r.Group("/query")
	query.POST("/search", controller.QueryByKeywords)
	query.GET("/model/:id", controller.QueryByModelID)
}
