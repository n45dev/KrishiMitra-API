package main

import (
	"krishimitra-api/db"
	"krishimitra-api/routes"
)

func main() {
	db.ConnectDB()
	router := routes.SetupRouter()
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
