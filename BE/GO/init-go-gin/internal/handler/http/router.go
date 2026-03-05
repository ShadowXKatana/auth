package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, userHandler *UserHandler) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(nethttp.StatusOK, gin.H{"status": "ok"})
	})

	v1 := engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.List)
		}
	}
}
