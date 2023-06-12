package model

import (
	"product-api/config"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(config.JWTSecret)

type Claims struct {
	UserID             int    `json:"user_id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Scope              string `json:"scope"`
	jwt.StandardClaims        // struct embedding
}
