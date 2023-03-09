package middleware

import (
	"go-example-api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
}

func AuthRequired(c *gin.Context) {
	access_token, err := c.Cookie("access_token")

	if err != nil {
		authError(c)
		return
	}

	_, status := helpers.ParseAccessToken(access_token)

	if !status {
		authError(c)
		return
	}

	c.Next()
}
