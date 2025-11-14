package handler

import (
	"net/http"

	"myapp/internal/infrastructure/logger"
	"myapp/pkg/errors"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	logger  logger.Logger
	errors.AppError
}

func NewHealthCheckHandler(logger logger.Logger) *HealthCheckHandler {
	return &HealthCheckHandler{AppError: *errors.New(logger), logger: logger}
}

func (h *HealthCheckHandler) GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
		"status":  "healthy",
	})
}