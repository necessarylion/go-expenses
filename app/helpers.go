package app

import (
	"fmt"
	"strconv"
)

func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func ToInt(value interface{}) int {
	v, _ := strconv.Atoi(ToString(value))
	return v
}
