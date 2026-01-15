package controllers

import (
	"net/http"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/dto"
	"github.com/SaraLuciaa/ssl-labs-challenge/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		"analysis": analysis,
	})
}

func (ctrl *AnalysisController) GetAnalysis(c *gin.Context) {
	idParam := c.Param("id")
	analysisID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid analysis ID format",
		})
		return
	}

	analysis, err := ctrl.service.GetAnalysisById(analysisID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"analysis": analysis,
	})
}
