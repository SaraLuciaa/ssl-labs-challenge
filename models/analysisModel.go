package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Analysis struct {
	gorm.Model
	ID              uuid.UUID      
	Host            string         
	Status          string         
	StatusMessage   string         
	Grade           string         
	StartTime       *time.Time
	EndTime         *time.Time     
	IsPublic        bool           
	EngineVersion   string         
	CriteriaVersion string         
	RawResponse     map[string]any 
	CreatedAt       time.Time      
	UpdatedAt       time.Time      
}