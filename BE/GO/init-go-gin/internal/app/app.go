package app

import (
	"github.com/gin-gonic/gin"
	userHTTP "github.com/sos/auth/be/go/init-go-gin/internal/delivery/http/handler/user"
	"github.com/sos/auth/be/go/init-go-gin/internal/delivery/http/router"
	userRepo "github.com/sos/auth/be/go/init-go-gin/internal/repository/memory/user"
	userUsecase "github.com/sos/auth/be/go/init-go-gin/internal/usecase/user"
)

func New() *gin.Engine {
	engine := gin.Default()

	repository := userRepo.NewRepository()
	usecase := userUsecase.NewUsecase(repository)
	handler := userHTTP.NewHandler(usecase)

	router.Register(engine, handler)

	return engine
}
