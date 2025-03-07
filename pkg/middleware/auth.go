package middleware

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireToken is a middleware function for Gin that checks for a valid
// Authorization token in the request header. If the token is missing or invalid,
// it responds with an Unauthorized status and aborts the request. Otherwise, it
// sets the token claims in the context and proceeds to the next handler.
func RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		if token == "" {
			c.JSON(http.StatusUnauthorized, &dto.ResponseDto{
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, &dto.ResponseDto{
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
