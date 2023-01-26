package gin_routes

import (
	"parishioner_management/internal/components"

	"github.com/gin-gonic/gin"

	auth_transport "parishioner_management/internal/modules/auth/transport"
	parishioner_transport "parishioner_management/internal/modules/parishioner/transport"
)

const (
	authPrefix         = "auth"
	parishionersPrefix = "parishioners"
)

const (
	authLoginSuffix   = "login"
	authRefreshSuffix = "refresh"
	listSuffix        = "list"
	newSuffix         = "new"
	deleteSuffix      = "delete"
)

func RegisterAllModules(g *gin.Engine, appContext components.AppContext, apiPrefix string) {
	db := appContext.GetMainMongoDBConnection()

	api := g.Group(apiPrefix)
	{

		auth := api.Group(authPrefix)
		{
			// POST /api/auth/login
			auth.POST(authLoginSuffix, auth_transport.Login(db))

			// POST /api/auth/refresh
			auth.POST(authRefreshSuffix, auth_transport.Refresh(db))
		}

		parishioner := api.Group(parishionersPrefix)
		{
			// POST /api/parishioners/list
			parishioner.GET(listSuffix, parishioner_transport.GetList(db))

			// POST /api/parishioners/new
			parishioner.POST(newSuffix, parishioner_transport.Create(db))

			// DELETE /api/parishioners/delete
			parishioner.DELETE(deleteSuffix, parishioner_transport.Delete(db))
		}
	}
}
