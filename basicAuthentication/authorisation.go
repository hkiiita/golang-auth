package basicAuthentication

import (
	"github.com/gin-gonic/gin"
	"github.com/hkbiet/golang-basic-auth/utils"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	credentials := utils.GetCreds()
	if credentials != nil {
		return gin.BasicAuth(credentials)
	}
	return nil
}
