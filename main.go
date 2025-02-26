package main

import (
	"go-auth-mvc/config"
	"go-auth-mvc/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()
	r.Run(":9292")
}
