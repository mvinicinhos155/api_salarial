package token

import (
	"api_salarial/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

	func GenerateJwt(user models.User) (string, error) {
		var jwtKey = []byte(os.Getenv("JWT_SECRET"))
		
		claims := models.Claims{
			ID: strconv.Itoa(user.ID),
			Email: user.Email,
			Role: user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(18 * time.Hour).Unix(),
				IssuedAt: time.Now().Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


		tokenString, err := token.SignedString(jwtKey)
			if  err != nil {
				return "", err
			}

			return tokenString, nil
	}