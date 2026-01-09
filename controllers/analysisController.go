package controllers

import (
	"net/http"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/dto"
	"github.com/SaraLuciaa/ssl-labs-challenge/services"
	"github.com/gin-gonic/gin"
)

type AnalysisController struct {
	service *services.AnalysisService
}

func NewAnalysisController(service *services.AnalysisService) *AnalysisController {
	return &AnalysisController{
		service: service,
	}
}

func (ctrl *AnalysisController) AnalysisStart(c *gin.Context) {
	var request dto.AnalysisRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	analysis, err := ctrl.service.StartAnalysis(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Analysis started",
		"analysis": analysis,
	})
}
