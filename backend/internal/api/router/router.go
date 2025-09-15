package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default()) // 允许跨域
	RegisterAuthRoutes(r)
	RegisterWikiRoutes(r)
	return r
}
