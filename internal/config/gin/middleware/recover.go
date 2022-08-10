package gin_middleware

import (
	"init_golang/internal/common"
	"init_golang/internal/components"

	"github.com/gin-gonic/gin"
)

func recoverMiddleware(appContext components.AppContext) gin.HandlerFunc {
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
