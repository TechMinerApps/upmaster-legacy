package v1

import (
	"github.com/TechMinerApps/upmaster/modules/oauth"
	oauthApi "github.com/TechMinerApps/upmaster/router/api/v1/oauth"
	"github.com/gin-gonic/gin"
)

type Config struct {
	OAuthConfig oauth.Config
}

func NewRouter(c Config) *gin.Engine {
	router := gin.Default()
	oauthServer, _ := oauth.NewServer(c.OAuthConfig)
	router.Any("/oauth", oauthApi.Wrapper(oauthServer))

	return router
}
