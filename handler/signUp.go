package handler

import (
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Repo db.UserRepository
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	// body params を構造体にbind
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	// password をハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
	}

	// user model
	user := db.User{
		Name:     body.Name,
		Password: string(hash),
	}

	if err := h.Repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "name": user.Name})
}
