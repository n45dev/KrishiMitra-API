package main

import (
	"github.com/joho/godotenv"
	"krishimitra-api/db"
	"krishimitra-api/routes"
	"os"
)

func main() {
	db.ConnectDB()
	router := routes.SetupRouter()

	loadEnvErr := godotenv.Load(".env")
	if loadEnvErr != nil {
		return
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	err := router.Run(host + ":" + port)
	if err != nil {
		return
	}
}
