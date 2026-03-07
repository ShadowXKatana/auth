package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gormdb "github.com/sos/auth/be/go/my-storage-service/internal/gorm"
	handler "github.com/sos/auth/be/go/my-storage-service/internal/handler/http"
	repository "github.com/sos/auth/be/go/my-storage-service/internal/repository"
	"github.com/sos/auth/be/go/my-storage-service/internal/usecase"
	"github.com/sos/auth/be/go/my-storage-service/pkg"
)

func main() {
	engine := gin.Default()
	engine.Use(handler.CORS())

	// Database
	db, err := gormdb.NewGormDB(getDBDSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := gormdb.AutoMigrate(db); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	// Security
	accessTokenService := pkg.NewJWTService(getJWTSecret(), getJWTAccessTTL())
	refreshTokenService := pkg.NewJWTService(getJWTRefreshSecret(), getJWTRefreshTTL())
	passwordService := pkg.NewPasswordService()

	// Repositories
	userRepo := repository.NewUserRepository(db)
	storageRepo := repository.NewStorageRepository(db)
	itemRepo := repository.NewItemRepository(db)

	// Usecases
	userUC := usecase.NewUserUsecase(userRepo, passwordService, accessTokenService, refreshTokenService)
	storageUC := usecase.NewStorageUsecase(storageRepo)
	itemUC := usecase.NewItemUsecase(itemRepo)

	// Seed default user
	seedDefaultUser(userUC)

	// Handlers
	authHandler := handler.NewAuthHandler(userUC, isCookieSecure())
	storageHandler := handler.NewStorageHandler(storageUC)
	itemHandler := handler.NewItemHandler(itemUC)
	authMiddleware := handler.JWTAuth(accessTokenService)

	// Routes
	handler.RegisterRoutes(engine, authHandler, storageHandler, itemHandler, authMiddleware)

	// Start
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// ── Helpers ─────────────────────────────────────────────────────────────────

func getDBDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn != "" {
		return dsn
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// If any required var is missing, return fallback or error
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Println("Database configuration missing in env. Falling back to local default...")
		return "host=localhost user=auth password=auth dbname=auth port=5432 sslmode=disable TimeZone=UTC"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)
}

func getJWTSecret() string {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		return "dev-secret-change-me"
	}
	return s
}

func getJWTRefreshSecret() string {
	s := os.Getenv("JWT_REFRESH_SECRET")
	if s == "" {
		return "dev-refresh-secret-change-me"
	}
	return s
}

func getJWTAccessTTL() time.Duration {
	raw := os.Getenv("JWT_EXPIRES_IN_MINUTES")
	if raw == "" {
		return 60 * time.Minute
	}
	minutes, err := strconv.Atoi(raw)
	if err != nil || minutes <= 0 {
		return 60 * time.Minute
	}
	return time.Duration(minutes) * time.Minute
}

func getJWTRefreshTTL() time.Duration {
	raw := os.Getenv("JWT_REFRESH_EXPIRES_IN_MINUTES")
	if raw == "" {
		return 7 * 24 * time.Hour
	}
	minutes, err := strconv.Atoi(raw)
	if err != nil || minutes <= 0 {
		return 7 * 24 * time.Hour
	}
	return time.Duration(minutes) * time.Minute
}

func isCookieSecure() bool {
	return os.Getenv("APP_ENV") == "production"
}

func seedDefaultUser(uc usecase.UserUsecase) {
	email := os.Getenv("AUTH_SEED_EMAIL")
	password := os.Getenv("AUTH_SEED_PASSWORD")
	if email == "" || password == "" {
		return
	}
	_, _ = uc.Register(context.Background(), usecase.RegisterInput{
		Email:    email,
		Password: password,
	})
}

// test
