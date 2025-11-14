package middlewares

import (
	"fmt"
	"myapp/internal/infrastructure/logger"
	"myapp/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MiddlewareManager struct{
	Logger logger.Logger
	errors.AppError
}

func New(logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{Logger: logger}
}

func (m *MiddlewareManager) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.Request.Context().Value("correlationId")
		var correlationIdStr string
		if cid, ok := correlationId.(uuid.UUID); ok {
			correlationIdStr = cid.String()
		} else if cid, ok := correlationId.(string); ok {
			correlationIdStr = cid
		} else {
			correlationIdStr = ""
		}

		t := time.Now()
		
		c.Next()
		
		latency := time.Since(t)
		m.Logger.PrintInfo(fmt.Sprintf("%s request at %s", c.Request.Method, c.Request.URL.Path), map[string]string{
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"host":          c.Request.Host,
			"correlationId": correlationIdStr,
			"latency":       latency.String(),
			"status":        fmt.Sprintf("%d", c.Writer.Status()),
		})

	}
}