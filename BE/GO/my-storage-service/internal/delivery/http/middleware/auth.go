package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/my-storage-service/internal/security"
)

const accessTokenCookieName = "access_token"

func JWTAuth(tokenService security.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(accessTokenCookieName)
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing auth cookie"})
			return
		}

		claims, err := tokenService.Parse(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		c.Set("auth_user_id", claims.UserID)
		c.Set("auth_user_email", claims.Email)
		c.Next()
	}
}

func JWTAuthOptional(tokenService security.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(accessTokenCookieName)
		if err != nil || token == "" {
			c.Next()
			return
		}

		claims, err := tokenService.Parse(token)
		if err == nil {
			c.Set("auth_user_id", claims.UserID)
			c.Set("auth_user_email", claims.Email)
		}

		c.Next()
	}
}
