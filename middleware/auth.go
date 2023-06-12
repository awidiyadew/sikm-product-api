package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"product-api/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func validateJWT(c *gin.Context) {
	tokenValue, err := c.Cookie("session_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
		return
	}

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return model.JwtKey, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
		return
	}

	// cast claims interface to mapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
		return
	}

	// convert map to json then turn it into struct model.Claims
	b, _ := json.Marshal(claims)
	var customClaims model.Claims
	json.Unmarshal(b, &customClaims)

	// set jwt payload to gin context that can be shared within a request
	c.Set("user_id", customClaims.UserID)
	c.Set("email", customClaims.Email)
	c.Set("isAdmin", customClaims.Scope == "ADMIN")
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateJWT(c)
		c.Next()
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateJWT(c)

		// get isAdmin data from context
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, model.NewErrorResponse("admin access required!"))
			return
		}

		c.Next()
	}
}
