package app

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	authHTTP "github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/handler/auth"
	graphqlHTTP "github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/handler/graphql"
	storageHTTP "github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/handler/storage"
	"github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/middleware"
	"github.com/sos/auth/be/go/my-storage-service/internal/delivery/http/router"
	storageRepo "github.com/sos/auth/be/go/my-storage-service/internal/repository/memory/storage"
	postgresRepo "github.com/sos/auth/be/go/my-storage-service/internal/repository/postgres"
	userRepo "github.com/sos/auth/be/go/my-storage-service/internal/repository/postgres/user"
	"github.com/sos/auth/be/go/my-storage-service/internal/security"
	storageUsecase "github.com/sos/auth/be/go/my-storage-service/internal/usecase/storage"
	userUsecase "github.com/sos/auth/be/go/my-storage-service/internal/usecase/user"
)

func New() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.CORS())

	accessTokenService := security.NewJWTService(getJWTSecret(), getJWTAccessTTL())
	refreshTokenService := security.NewJWTService(getJWTRefreshSecret(), getJWTRefreshTTL())

	storageRepository := storageRepo.NewRepository()
	storageUC := storageUsecase.NewUsecase(storageRepository)
	storageHandler := storageHTTP.NewHandler(storageUC)

	db, err := postgresRepo.NewGormDB(getDBDSN())
	if err != nil {
		panic(err)
	}

	userRepository := userRepo.NewRepository(db)
	passwordService := security.NewPasswordService()
	userUC := userUsecase.NewUsecase(userRepository, passwordService, accessTokenService, refreshTokenService)
	seedDefaultUser(userUC)
	authHandler := authHTTP.NewHandler(userUC, isCookieSecure())
	graphqlHandler, err := graphqlHTTP.NewHandler()
	if err != nil {
		panic(err)
	}
	authMiddleware := middleware.JWTAuth(accessTokenService)

	router.Register(engine, authHandler, graphqlHandler, storageHandler, authMiddleware)

	return engine
}

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "dev-secret-change-me"
	}

	return secret
}

func getDBDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return "host=localhost user=auth password=auth dbname=auth port=5432 sslmode=disable TimeZone=UTC"
	}

	return dsn
}

func getJWTAccessTTL() time.Duration {
	rawTTL := os.Getenv("JWT_EXPIRES_IN_MINUTES")
	if rawTTL == "" {
		return 60 * time.Minute
	}

	minutes, err := strconv.Atoi(rawTTL)
	if err != nil || minutes <= 0 {
		return 60 * time.Minute
	}

	return time.Duration(minutes) * time.Minute
}

func getJWTRefreshSecret() string {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		return "dev-refresh-secret-change-me"
	}

	return secret
}

func getJWTRefreshTTL() time.Duration {
	rawTTL := os.Getenv("JWT_REFRESH_EXPIRES_IN_MINUTES")
	if rawTTL == "" {
		return 7 * 24 * time.Hour
	}

	minutes, err := strconv.Atoi(rawTTL)
	if err != nil || minutes <= 0 {
		return 7 * 24 * time.Hour
	}

	return time.Duration(minutes) * time.Minute
}

func isCookieSecure() bool {
	return os.Getenv("APP_ENV") == "production"
}

func seedDefaultUser(authUsecase userUsecase.Usecase) {
	email := os.Getenv("AUTH_SEED_EMAIL")
	password := os.Getenv("AUTH_SEED_PASSWORD")
	if email == "" || password == "" {
		return
	}

	_, _ = authUsecase.Register(context.Background(), userUsecase.RegisterInput{
		Email:    email,
		Password: password,
	})
}
