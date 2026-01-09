package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Analysis struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Host            string
	Status          string
	StatusMessage   string
	Grade           string
	StartTime       *time.Time
	EndTime         *time.Time
	IsPublic        bool
	EngineVersion   string
	CriteriaVersion string
	RawResponse     datatypes.JSON
	CreatedAt       time.Time
	UpdatedAt       time.Time
}