package controllers

import (
	"expenses/app"
	"expenses/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetRecords(c echo.Context) error {

	// Parse the request body into a struct
	type requestBody struct {
		Limit int `json:"limit" validate:"required"`
		Page  int `json:"page" validate:"required"`
	}
	limit, _ := strconv.Atoi(c.FormValue("limit"))
	page, _ := strconv.Atoi(c.FormValue("page"))

	u := requestBody{
		Limit: limit,
		Page:  page,
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	var count int64
	app.DB.Model(&models.User{}).Count(&count)

	offset := (u.Page - 1) * u.Limit
	users := []models.User{}
	result := app.DB.Limit(u.Limit).Offset(offset).Find(&users)

	if result.Error != nil {
		return c.String(http.StatusOK, "something wrong")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total":        count,
		"current_page": u.Page,
		"data":         users,
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
	user := models.User{}
	result := app.DB.Find(&user, c.Param("id"))

	if result.Error != nil {
		return c.String(http.StatusOK, "something wrong")
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteRecord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})
}
