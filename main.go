package main

import (
	"expenses/app"
	"expenses/routes"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	app.Env()
	app.DatabaseConnection()
	RegisterRoutes()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func RegisterRoutes() {
	var e *echo.Echo = echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	routes.RegisterWebRoutes(e)
	routes.RegisterApiRoutes(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
