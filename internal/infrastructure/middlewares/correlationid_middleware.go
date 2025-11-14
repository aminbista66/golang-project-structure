package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *MiddlewareManager) AssignCorrelationId() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId, err := uuid.NewUUID()

		if err != nil {
			m.Logger.PrintError("failed to create correlation ID", map[string]string{"error": err.Error()})
			m.ServerErrorResponse(c.Writer, c.Request, err)
		}

		ctx := c.Request.Context()
		newCtxWithValue := context.WithValue(ctx, "correlationId", correlationId)
		newCtxWithTimeout, cancel := context.WithTimeout(newCtxWithValue, time.Minute)
		defer cancel()

		newRequest := c.Request.WithContext(newCtxWithTimeout)

		c.Request = newRequest
		c.Next()
	}
}