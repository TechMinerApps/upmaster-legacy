package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RegisterLogger register a logger into gin context
func RegisterLogger(l *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Logger", l)
		c.Next()
	}
}
