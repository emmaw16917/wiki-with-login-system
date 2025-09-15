package handler

import (
	"net/http"
	"wiki/backend/internal/model"
	"wiki/backend/internal/service"
	"wiki/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) {
	slug := c.Param("slug")
	page, err := service.GetPageBySlug(database.DB, slug)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "页面不存在"})
		return
	}
	c.IndentedJSON(http.StatusOK, page)
}

func CreatePage(c *gin.Context) {
	var page model.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	userID, _ := c.Get("user_id")
	id, err := service.CreatePage(database.DB, &page, userID.(int))
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"page_id": id})
}

func UpdatePage(c *gin.Context) {
	slug := c.Param("slug")
	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "content字段解析错误"})
		return
	}
	userID, _ := c.Get("user_id")
	page, err := service.GetPageBySlug(database.DB, slug)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "页面不存在"})
		return
	}

	err = service.UpdatePageContent(database.DB, page.ID, req.Content, userID.(int))
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "页面已更新"})
}

func DeletePage(c *gin.Context) {
	slug := c.Param("slug")
	userID, _ := c.Get("user_id")
	page, err := service.GetPageBySlug(database.DB, slug)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "页面不存在"})
		return
	}

	err = service.DeletePage(database.DB, page.ID, userID.(int))
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "页面已删除"})
}
