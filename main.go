package main

import (
	"learn-gin-gorm/configs"
	"learn-gin-gorm/routes"
)

func main() {
	configs.InitDB()

	router := routes.AuthRouter()

	router.Run(":8080")
}
