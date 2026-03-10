package models

import "github.com/golang-jwt/jwt"

type Claims struct{
		ID string `json:"id"`
		Email string `json:"email"`
		Role string `json:"role"`
		jwt.StandardClaims
}