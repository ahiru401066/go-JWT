package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	// Context から userID を取得
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "userID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to your dash board",
		"userID":  userID,
	})
}
