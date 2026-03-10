package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"api_salarial/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var claims = models.Claims{}

func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Authorization errada ou invalida"})
			c.Abort()
			return 
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		    if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		        return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		    }
		    return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalido ou expirou token"})
			c.Abort()
			return 
		}

		c.Set("user", claims)
		c.Next()

	}
}


func AdminOnly() gin.HandlerFunc{
	return func(c *gin.Context) {
		user, exists := c.Get("user")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario não logado"})
				c.Abort()
				return 
			}

		claims, ok := user.(*models.Claims)
		if !ok || claims.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "acesso negado"})
			c.Abort()
			return 
		}

		c.Next()
	}
}