package main

import (
	"nu/corpus-reader/infrastructure"
	"nu/corpus-reader/infrastructure/router"
)

func main() {
	var app = infrastructure.NewConfig().Logger()
	app.WebServerPort("8080").WebServer(router.Gin).Start()
}
