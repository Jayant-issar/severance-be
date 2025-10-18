package middleware

import (
	"net/http"
	"strings"

	"github.com/Jayant-issar/severance-backend/severance-api/internal/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			util.SendResponse(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := util.ValidateJWT(token)
		if err != nil {
			util.SendResponse(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
