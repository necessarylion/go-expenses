package controllers

import (
	"expenses/app"
	"expenses/models"
	request "expenses/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRecords(c echo.Context) error {
	reqMap := request.All(c)

	var req request.ListingRequest
	request.Bind(reqMap, &req)

	if err := c.Validate(req); err != nil {
		return err
	}

	limit := app.ToInt(req.Limit)
	page := app.ToInt(req.Page)

	var count int64
	app.DB.Model(&models.User{}).Count(&count)

	offset := (page - 1) * limit
	users := []models.User{}
	result := app.DB.Limit(limit).Offset(offset).Find(&users)

	if result.Error != nil {
		return c.String(http.StatusOK, "something wrong")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total":        count,
		"current_page": page,
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
