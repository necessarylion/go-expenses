package request

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func All(c echo.Context) map[string]interface{} {
	query := c.QueryParams()

	jsonMap := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&jsonMap)

	for index, element := range query {
		jsonMap[index] = element[0]
	}
	return jsonMap
}

func Bind(m map[string]interface{}, s interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
