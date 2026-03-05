package app

import (
	"github.com/gin-gonic/gin"
	handler "github.com/sos/auth/be/go/init-go-gin/internal/handler/http"
	"github.com/sos/auth/be/go/init-go-gin/internal/repository"
	"github.com/sos/auth/be/go/init-go-gin/internal/usecase"
)

func New() *gin.Engine {
	engine := gin.Default()

	// Repositories
	userRepo := repository.NewUserRepository()

	// Usecases
	userUC := usecase.NewUserUsecase(userRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userUC)

	// Routes
	handler.RegisterRoutes(engine, userHandler)

	return engine
}
