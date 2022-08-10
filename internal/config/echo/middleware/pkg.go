package echo_middleware

import (
	"init_golang/internal/components"
	options_util "init_golang/internal/utils/options"

	"github.com/labstack/echo/v4"
)

//This file will contains all middleware registration for echo, simplify main server config
func RegisterMiddleware(e *echo.Echo, appContext components.AppContext, options options_util.Options) error {
	return nil
}
