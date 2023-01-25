package gin_routes

import (
	"parishioner_management/internal/components"

	"github.com/gin-gonic/gin"

	auth_transport "parishioner_management/internal/modules/auth/transport"
)

const (
	authPrefix        = "auth"
	authLoginSuffix   = "login"
	authRefreshSuffix = "refresh"
)

func RegisterAllModules(g *gin.Engine, appContext components.AppContext, apiPrefix string) {
	db := appContext.GetMainMongoDBConnection()

	api := g.Group(apiPrefix)
	{
		// POST /api/auth/login
		auth := api.Group(authPrefix)
		{
			auth.POST(authLoginSuffix, auth_transport.Login(db))
			auth.POST(authRefreshSuffix, auth_transport.Refresh(db))
		}
	}
}
