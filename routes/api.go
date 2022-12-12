package routes

import (
	"expenses/controllers"
	"expenses/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterApiRoutes(e *echo.Echo) {
	Route := e.Group("/api")
	Route.Use(middleware.AuthMiddleware)

	Route.POST("/auth/register", controllers.RegisterNewUser)
	Route.POST("/auth/login", controllers.LoginUser)

	Route.POST("/records", controllers.CreateRecord)
	Route.GET("/records", controllers.GetRecords)
	Route.GET("/records/:id", controllers.GetRecord)
	Route.PUT("/records/:id", controllers.UpdateRecord)
	Route.DELETE("/records/:id", controllers.DeleteRecord)
}
