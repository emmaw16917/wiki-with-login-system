package router

import (
	"wiki/backend/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}
