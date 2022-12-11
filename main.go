package main

import (
	"expenses/start"
)

func init() {
	start.Env()
	start.DatabaseConnection()
}

func main() {
	start.Server()
}
