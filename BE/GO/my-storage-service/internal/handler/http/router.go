package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	engine *gin.Engine,
	auth *AuthHandler,
	storage *StorageHandler,
	item *ItemHandler,
	authMiddleware gin.HandlerFunc,
) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(nethttp.StatusOK, gin.H{"status": "ok"})
	})

	v1 := engine.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", auth.Register)
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/refresh", auth.Refresh)
			authGroup.POST("/logout", auth.Logout)
			authGroup.GET("/me", authMiddleware, auth.Me)
		}

		storages := v1.Group("/storages", authMiddleware)
		{
			storages.POST("", storage.Create)
			storages.GET("", storage.List)
			storages.GET("/:id", storage.GetByID)
			storages.DELETE("/:id", storage.Delete)

			storages.POST("/:id/items", item.Create)
			storages.GET("/:id/items", item.List)
			storages.DELETE("/:id/items/:itemId", item.Delete)
			storages.PATCH("/:id/items/:itemId/tags", item.UpdateTags)
		}
	}
}
