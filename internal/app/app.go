package app

import (
	"os"
	"sync"

	"myapp/internal/config"
	"myapp/internal/domain/user"
	userpg "myapp/internal/domain/user"
	"myapp/internal/infrastructure/database"
	"myapp/internal/infrastructure/http/handler"
	"myapp/internal/infrastructure/logger"
	jsonlog "myapp/internal/infrastructure/logger/jsonlog"

	"github.com/gin-gonic/gin"
)

var L = jsonlog.New(os.Stdout, jsonlog.LevelInfo)

type Container struct {
	Router *gin.Engine
	Logger logger.Logger
	wg sync.WaitGroup
}

func Initialize(cfg *config.Config) (*Container, error) {
	db, err := database.NewPostgres(cfg.Database.DSN)
	if err != nil {
		return nil, err
	}

	userRepo := userpg.NewPostgresRepository(db)
	userSvc := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userSvc, L)

	router := gin.Default()
	router.GET("/api/users/:id", userHandler.GetUser)
	router.POST("/api/users", userHandler.CreateUser)

	return &Container{Router: router, Logger: L}, nil
}
