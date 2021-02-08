package v1

import (
	"fmt"

	"github.com/TechMinerApps/upmaster/modules/oauth"
	"github.com/gin-gonic/gin"
)

type Config struct {
	OAuthConfig oauth.Config
}

func NewRouter(c Config) (*gin.Engine, error) {
	router := gin.Default()

	oauthSubrouter := router.Group("/oauth/")
	oauthAPI, err := oauth.NewServer(c.OAuthConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to create OAuth server, %v", err)
	}
	oauthAPI.RegisterRoute(oauthSubrouter)

	return router, nil
}
