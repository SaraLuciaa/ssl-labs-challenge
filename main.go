package main

import (
	"github.com/SaraLuciaa/ssl-labs-challenge/controllers"
	"github.com/SaraLuciaa/ssl-labs-challenge/initializers"
	"github.com/SaraLuciaa/ssl-labs-challenge/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.DB.AutoMigrate(&models.Analysis{})
}

func main() {
	r := gin.Default()

	r.POST("/analysis", controllers.AnalysisStart)

	r.Run()
}
