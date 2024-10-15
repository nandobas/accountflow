package handlers

import (
	"accountflow/api/middlewares"
	"accountflow/modules/auth"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	//logica de login

	token, err := auth.NewJWTAuth("my-secret-key").GenerateToken("0002")
	if err != nil {
		Response(c, middlewares.RetFail(fmt.Sprintf("unable to generate token: %s", err)))
		return
	}

	Response(c, middlewares.RetOkData(token))
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenizer := auth.NewJWTAuth("my-secret-key")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := tokenizer.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
