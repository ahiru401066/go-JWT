package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *gin.Context) {
	// Cookie から JWT を取得
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no token"})
		return
	}

	// token をパース & 検証
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	// 有効確認
	if err := claims.Valid(); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token expired or invalid"})
		return
	}

	// userID を Context に保存
	c.Set("userID", claims.Subject)

	c.Next()
}
