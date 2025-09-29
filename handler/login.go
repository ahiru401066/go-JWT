package handler

import (
	"errors"
	"fmt"
	"main/db"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginHandler struct {
	Repo db.UserRepository
}

func (h *LoginHandler) Login(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	// body params を bind
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	// DBから、userNameでuser情報取得
	user, err := h.Repo.FindByUserName(body.Name)

	// user not foundの処理
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "UserNotFound",
		})
		return
	}

	// error処理（user not found 以外の処理）
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalServerError",
		})
		return
	}

	// password check
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The password is incorrect",
		})
		return
	}

	// JWT の作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", user.ID), // 文字列化
		// 有効期限設定
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// 秘密鍵で署名
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
