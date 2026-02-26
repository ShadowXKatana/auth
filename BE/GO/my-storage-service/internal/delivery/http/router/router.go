package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authHTTP "github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/handler/auth"
	graphqlHTTP "github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/handler/graphql"
	"github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/handler/storage"
)

func Register(engine *gin.Engine, authHandler *authHTTP.Handler, graphqlHandler *graphqlHTTP.Handler, storageHandler *storage.Handler, authMiddleware gin.HandlerFunc) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	engine.GET("/graphql", graphqlHandler.Serve)
	engine.POST("/graphql", graphqlHandler.Serve)

	v1 := engine.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.Refresh)
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", authMiddleware, authHandler.Me)
		}

		storages := v1.Group("/storages", authMiddleware)
		{
			storages.POST("", storageHandler.Create)
			storages.GET("", storageHandler.List)
		}
	}
}
