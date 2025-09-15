package router

import (
	"wiki/backend/internal/api/handler"
	"wiki/backend/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterWikiRoutes(r *gin.Engine) {
	r.GET("/pages/:slug", handler.GetPage)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/pages", handler.CreatePage)
		auth.PUT("/pages/:slug", handler.UpdatePage)
		auth.DELETE("/pages/:slug", handler.DeletePage)
	}
}
