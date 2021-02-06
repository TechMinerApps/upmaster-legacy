package router

import (
	v1 "github.com/TechMinerApps/upmaster/router/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Config has basic configurations
// Detailed config should be generated from basic config by database query
type Config struct {
	DB              *gorm.DB
	DBName          string
	OAuthGCInterval int
}

// NewRouter generates a gin router from config
func NewRouter(c Config) *gin.Engine {

	// Init a api configuration
	var apiConfig v1.Config
	apiConfig.OAuthConfig.DB = c.DB
	apiConfig.OAuthConfig.DBName = c.DBName
	apiConfig.OAuthConfig.GCInterval = c.OAuthGCInterval
	apiConfig.OAuthConfig.JWTKey = []byte("") // Wait for implementation
	apiConfig.OAuthConfig.Clients = nil       // Same

	return v1.NewRouter(apiConfig)
}
