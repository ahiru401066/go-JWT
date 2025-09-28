package handler

import (
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo db.UserRepository
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	// user model
	user := db.User{
		Name:     body.Name,
		Password: body.Password,
	}

	if err := h.Repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "name": user.Name})
}
