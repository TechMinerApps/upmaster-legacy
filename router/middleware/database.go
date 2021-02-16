package middleware

import (
	"github.com/TechMinerApps/upmaster/modules/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRDBMSInstance register DB to gin context
func RegisterRDBMSInstance(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

// RegisterInfluxDBInstance register InfluxDB to gin context
func RegisterInfluxDBInstance(i *database.InfluxDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("InfluxDB", i)
		c.Next()
	}
}
