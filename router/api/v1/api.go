package v1

import (
	"fmt"

	"github.com/TechMinerApps/upmaster/modules/oauth"
	"github.com/TechMinerApps/upmaster/router/api/v1/auth"
	"github.com/TechMinerApps/upmaster/router/api/v1/endpoints"
	"github.com/TechMinerApps/upmaster/router/api/v1/status"
	"github.com/gin-gonic/gin"
)

type Config struct {
	OAuthConfig oauth.Config
}

func SetupRouter(c Config, router *gin.Engine) error {

	oauthSubrouter := router.Group("/oauth/")
	oauthAPI, err := oauth.NewServer(c.OAuthConfig)
	if err != nil {
		return fmt.Errorf("Unable to create OAuth server, %v", err)
	}
	oauthAPI.RegisterRoute(oauthSubrouter)

	router.POST("/status", status.WriteEndpointStatus)

	authAPI := router.Group("auth")
	{
		authAPI.POST("/login", auth.Login)
		authAPI.DELETE("/logout", auth.Logout)
		authAPI.POST("/reset", auth.SendResetToken)
		authAPI.PUT("/reset", auth.ResetPassword)
	}

	endpointsAPI := router.Group("endpoints")
	{
		endpointsAPI.GET("", endpoints.Index)
		endpointsAPI.POST("", endpoints.Store)
		endpointsAPI.PUT("/:id", endpoints.Update)
		endpointsAPI.DELETE("/:id", endpoints.Destroy)
	}

	usersAPI := router.Group("users")
	{
		usersAPI.GET("", users.Index)
		usersAPI.PUT("/:username", users.Update)
		usersAPI.DELETE("/:username", users.Destroy)
	}

	return nil
}
