package main

import (
	"go-auth-mvc/config"
	"go-auth-mvc/routes"
)

func main() {
	//Setup Loging
	config.SetupLogger()

	//Logging saat aplikasi dimulai
	config.Log.Info("Starting the application...")

	config.ConnectDatabase()
	config.Log.Info("Database connected successfully")

	r := routes.SetupRouter()
	config.Log.Info("Server running on port 9292")
	r.Run(":9292")
}
