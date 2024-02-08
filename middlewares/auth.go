package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the authorization token from the request header
		token := context.GetHeader("Authorization")

		// Check if the token is empty
		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is required"})
			return
		}

		// Verify the token
		userId, err := utils.VerifyToken(token)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization token"})
			return
		}

		// Token is valid, continue processing
		context.Set("userId", userId)
		context.Next()
	}
}
