package gin_middleware

import (
	"parishioner_management/internal/common"
	"parishioner_management/internal/components"
	options_util "parishioner_management/internal/utils/options"

	"github.com/gin-gonic/gin"
)

func customRecover(appContext components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// panic(err)
				return
			}
		}()

		c.Next()
	}
}

func recoverMiddleware(g *gin.Engine, appContext components.AppContext, options options_util.Options) {
	g.Use(customRecover(appContext))
}
