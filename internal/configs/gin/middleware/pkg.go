package gin_middleware

import (
	"init_golang/internal/components"
	options_util "init_golang/internal/utils/options"

	"github.com/gin-gonic/gin"
)

//This file will contains all middleware registration for echo, simplify main server config
func RegisterMiddleware(g *gin.Engine, appContext components.AppContext, options options_util.Options) error {
	g.Use(recoverMiddleware(appContext))
	g.Use(gin.Logger())

	return nil
}
