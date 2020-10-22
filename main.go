package main

import (
	"go-with-compose/app"
	"go-with-compose/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":9567")
}
