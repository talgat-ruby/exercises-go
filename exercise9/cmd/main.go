package main

import (
	"alinurmyrzakhanov/database"
	"alinurmyrzakhanov/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	database.SeedExercises()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
