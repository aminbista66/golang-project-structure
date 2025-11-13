package handler

import (
	"net/http"
	"strconv"

	"myapp/internal/domain/user"
	"myapp/internal/infrastructure/logger"
	"myapp/pkg/errors"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *user.Service
	logger  logger.Logger
	errors.AppError
}

func NewUserHandler(service *user.Service, logger logger.Logger) *UserHandler {
	return &UserHandler{service: service, AppError: *errors.New(logger), logger: logger}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u, err := h.service.GetUser(id)
	if err != nil {
		h.NotFoundResponse(c.Writer, c.Request)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}
