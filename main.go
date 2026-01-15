package main

import (
	"net/http"

	"github.com/SaraLuciaa/ssl-labs-challenge/controllers"
	"github.com/SaraLuciaa/ssl-labs-challenge/initializers"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
	"github.com/SaraLuciaa/ssl-labs-challenge/repositories"
	"github.com/SaraLuciaa/ssl-labs-challenge/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.DB.AutoMigrate(&models.Analysis{}, &models.Endpoint{})
}

func main() {
	r := gin.Default()

	httpClient := &http.Client{}
	sslService := services.NewSSLLabsService(httpClient)
	analysisRepo := repositories.NewAnalysisRepository(initializers.DB)
	endpointRepo := repositories.NewEndpointRepository(initializers.DB)
	analysisService := services.NewAnalysisService(sslService, analysisRepo, endpointRepo)
	analysisController := controllers.NewAnalysisController(analysisService)

	r.POST("/analysis", analysisController.AnalysisStart)
	r.GET("/analysis/:id", analysisController.GetAnalysis)

	r.Run()
}
