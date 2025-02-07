package main

import (
	"alinurmyrzakhanov/database"
	"alinurmyrzakhanov/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
