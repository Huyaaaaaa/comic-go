package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte

func SetJWTSecret(secret string) {
	JWTSecret = []byte(secret)
}

type Claims struct {
	UserID  uint   `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, isAdmin bool) (string, error) {
	claims := Claims{
		UserID:  userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 无效"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("is_admin", claims.IsAdmin)
		c.Next()
	}
}

func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth != "" && strings.HasPrefix(auth, "Bearer ") {
			tokenStr := strings.TrimPrefix(auth, "Bearer ")
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
				return JWTSecret, nil
			})
			if err == nil && token.Valid {
				c.Set("user_id", claims.UserID)
				c.Set("is_admin", claims.IsAdmin)
			}
		}
		c.Next()
	}
}
