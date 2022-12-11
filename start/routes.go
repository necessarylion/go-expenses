package start

import (
	"expenses/controllers"
	"expenses/middleware"
	"os"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes() {
	var Route *echo.Echo = echo.New()
	RegisterWebRoutes(Route)
	RegisterApiRoutes(Route)
	Route.Start(":" + os.Getenv("APP_PORT"))
}

func RegisterWebRoutes(Route *echo.Echo) {
	Route.GET("/health", controllers.Ping)
}

func RegisterApiRoutes(e *echo.Echo) {
	Route := e.Group("/api")
	Route.Use(middleware.AuthMiddleware)

	Route.POST("/records", controllers.CreateRecord)
	Route.GET("/records", controllers.GetRecords)
	Route.GET("/records/:id", controllers.GetRecord)
	Route.PUT("/records/:id", controllers.UpdateRecord)
	Route.DELETE("/records/:id", controllers.DeleteRecord)
}
