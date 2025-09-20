package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

func NewUser(c *gin.Context) {
	var request Request
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": request.Name + request.PassWord,
	})
}
