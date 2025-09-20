package handler

import (
	"main/db"
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
		return
	}

	user := db.User{
		Name:     request.Name,
		Password: request.PassWord,
	}

	if err := db.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "name": user.Name})
}
