package controllers

import (
	request "expenses/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterNewUser(c echo.Context) error {
	reqMap := request.All(c)
	type Request struct {
		Email                string `json:"email" validate:"required,email"`
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required"`
		Name                 string `json:"name" validate:"required"`
	}
	var req Request
	request.Bind(reqMap, &req)

	if err := c.Validate(req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    req.Email,
	})
}

func LoginUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
