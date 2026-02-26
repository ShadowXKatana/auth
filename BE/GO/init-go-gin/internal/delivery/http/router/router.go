package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/init-go-gin/internal/delivery/http/handler/user"
)

func Register(engine *gin.Engine, userHandler *user.Handler) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
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
