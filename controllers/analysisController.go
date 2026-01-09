package controllers

import (
	"github.com/gin-gonic/gin"
)

func AnalysisStart (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Analysis started",
	})
}