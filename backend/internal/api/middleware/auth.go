package middleware

import (
	"net/http"
	"strings"
	"wiki/backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc { // JWT认证中间件
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "请求头未提供有效的Token"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Token无效或已过期"})
			c.Abort()
			return
		}

		c.Set("user_id", userID) //将解析出来的用户ID存入Gin上下文
		c.Next()
	}
}
