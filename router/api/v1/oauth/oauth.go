package oauth

import (
	oauthMod "github.com/TechMinerApps/upmaster/modules/oauth"
	"github.com/gin-gonic/gin"
)

func Wrapper(s *oauthMod.Server) gin.HandlerFunc {
	return gin.WrapH(s.Handler)
}
