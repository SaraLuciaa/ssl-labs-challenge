package main

import (
	"net/http"
	"os"

	"github.com/SaraLuciaa/ssl-labs-challenge/controllers"
	"github.com/SaraLuciaa/ssl-labs-challenge/initializers"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
	"github.com/SaraLuciaa/ssl-labs-challenge/repositories"
	"github.com/SaraLuciaa/ssl-labs-challenge/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.DB.AutoMigrate(&models.Analysis{}, &models.Endpoint{})
}

func main() {
	r := gin.Default()

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	httpClient := &http.Client{}
	sslService := services.NewSSLLabsService(httpClient)
	analysisRepo := repositories.NewAnalysisRepository(initializers.DB)
	endpointRepo := repositories.NewEndpointRepository(initializers.DB)
	analysisService := services.NewAnalysisService(sslService, analysisRepo, endpointRepo)
	analysisController := controllers.NewAnalysisController(analysisService)

	r.POST("/analysis", analysisController.AnalysisStart)
	r.GET("/analysis", analysisController.GetAllAnalyses)
	r.GET("/analysis/:id", analysisController.GetAnalysis)
	r.GET("/analysis/:id/location", analysisController.GetLocationById)

	r.Run()
}
