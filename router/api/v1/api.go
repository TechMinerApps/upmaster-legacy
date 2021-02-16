package v1

import (
	"fmt"

	"github.com/TechMinerApps/upmaster/modules/oauth"
	"github.com/TechMinerApps/upmaster/router/api/v1/status"
	"github.com/gin-gonic/gin"
)

type Config struct {
	OAuthConfig oauth.Config
}

func NewRouter(c Config) (*gin.Engine, error) {
	router := gin.Default()

	router.Static("/swagger/", "./docs")

	oauthSubrouter := router.Group("/oauth/")
	oauthAPI, err := oauth.NewServer(c.OAuthConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to create OAuth server, %v", err)
	}
	oauthAPI.RegisterRoute(oauthSubrouter)

	statusAPI := router.Group("/status/")
	statusAPI.POST("/:endpoint_id/", status.WriteEndpointStatus)

	return router, nil
}
