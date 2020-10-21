package main

import "go-with-compose/app"

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":9567")
}
