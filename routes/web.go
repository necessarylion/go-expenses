package routes

import (
	"expenses/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterWebRoutes(Route *echo.Echo) {
	Route.GET("/health", controllers.Ping)
}
