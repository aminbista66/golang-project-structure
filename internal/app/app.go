package app

import (
	"os"
	"sync"

	"myapp/internal/config"
	// "myapp/internal/domain/user"
	// userpg "myapp/internal/domain/user"
	// "myapp/internal/infrastructure/database"
	"myapp/internal/infrastructure/http/handler"
	"myapp/internal/infrastructure/logger"
	jsonlog "myapp/internal/infrastructure/logger/jsonlog"
	"myapp/internal/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
)


type Container struct {
	Router *gin.Engine
	Logger logger.Logger
	wg sync.WaitGroup
}

func Initialize(cfg *config.Config) (*Container, error) {
	// db, err := database.NewPostgres(cfg.Database.DSN)
	// if err != nil {
	// 	return nil, err
	// }

	// userRepo := userpg.NewPostgresRepository(db)
	// userSvc := user.NewService(userRepo)
	// userHandler := handler.NewUserHandler(userSvc, L)
	
	router := gin.New()
	container := &Container{Router: router, Logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo)}

	
	middlewares := middlewares.New(container.Logger)
	router.Use(gin.Recovery())
	router.Use(middlewares.AssignCorrelationId())
	router.Use(middlewares.LogRequest())

	router.GET("/api/health", handler.NewHealthCheckHandler(container.Logger).GetHealthCheck)

	// router.GET("/api/users/:id", userHandler.GetUser)
	// router.POST("/api/users", userHandler.CreateUser)

	return container, nil
}
