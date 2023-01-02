package gin_middleware

import (
	"parishioner_management/internal/components"
	options_util "parishioner_management/internal/utils/options"

	"github.com/gin-gonic/gin"
)

func loggerMiddleware(g *gin.Engine, appContext components.AppContext, options options_util.Options) {
	g.Use(gin.Logger())
}
