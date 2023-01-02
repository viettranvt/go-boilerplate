package gin_middleware

import (
	"parishioner_management/internal/components"
	options_util "parishioner_management/internal/utils/options"

	"github.com/gin-gonic/gin"
)

//This file will contains all middleware registration for echo, simplify main server config
func RegisterMiddleware(g *gin.Engine, appContext components.AppContext, options options_util.Options) error {
	recoverMiddleware(g, appContext, options)
	loggerMiddleware(g, appContext, options)
	secureMiddleware(g, appContext, options)

	return nil
}
