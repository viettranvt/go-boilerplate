package gin_middleware

import (
	"net/http"
	"os"
	"parishioner_management/internal/components"
	configs "parishioner_management/internal/configs"
	options_util "parishioner_management/internal/utils/options"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func secureMiddleware(g *gin.Engine, appCtx components.AppContext, options options_util.Options) {
	origin := os.Getenv(configs.EnvCorsOrigin)
	config := cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowOrigins:     strings.Split(origin, ","),
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	})

	g.Use(config)
}
