package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRecords(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})
}

func CreateRecord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})
}

func UpdateRecord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})
}

func GetRecord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})
}

func DeleteRecord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})
}
