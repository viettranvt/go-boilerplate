package gin_middleware

import (
	"net/http"
	"os"
	"parishioner_management/internal/components"
	options_util "parishioner_management/internal/utils/options"
	"strings"

	configs "parishioner_management/internal/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func secureMiddleware(g *gin.Engine, appCtx components.AppContext, options options_util.Options) {

	origin := os.Getenv(configs.EnvCorsOrigin)

	g.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     strings.Split(origin, ","),
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
}
