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

func SetupRouter(c Config, router *gin.Engine) error {

	oauthSubrouter := router.Group("/oauth/")
	oauthAPI, err := oauth.NewServer(c.OAuthConfig)
	if err != nil {
		return fmt.Errorf("Unable to create OAuth server, %v", err)
	}
	oauthAPI.RegisterRoute(oauthSubrouter)

	statusAPI := router.Group("/status/")
	statusAPI.POST("/:endpoint_id/", status.WriteEndpointStatus)

	return nil
}
