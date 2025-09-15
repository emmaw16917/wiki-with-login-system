package handler

import (
	"net/http"
	"wiki/backend/internal/service"
	"wiki/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	userID, err := service.RegisterUser(database.DB, req.Username, req.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"user_id": userID})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	user, err := service.LoginUser(database.DB, req.Username, req.Password)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := service.GenerateToken(user.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "生成Token失败"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
