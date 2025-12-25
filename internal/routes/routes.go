package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	v1api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(v1api)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"path":  ctx.Request.URL.Path,
		})
	})
}
