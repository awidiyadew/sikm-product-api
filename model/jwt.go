package model

import "github.com/golang-jwt/jwt"

var JwtKey = []byte("secret-key")

type Claims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Scope  string `json:"scope"`
	jwt.StandardClaims // struct embedding
}
