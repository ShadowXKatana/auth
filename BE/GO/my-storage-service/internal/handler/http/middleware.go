package http

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/my-storage-service/pkg"
)

// ── CORS Middleware ─────────────────────────────────────────────────────────

func CORS() gin.HandlerFunc {
	allowedOrigin := os.Getenv("APP_CORS_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000"
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" && origin == allowedOrigin {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Vary", "Origin")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// ── JWT Auth Middleware ─────────────────────────────────────────────────────

func JWTAuth(tokenService pkg.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
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
