package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Analysis struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Host            string         `gorm:"not null" json:"host"`
	Status          string         `gorm:"not null" json:"status"`
	StatusMessage   string         `json:"status_message,omitempty"`
	Grade           string         `json:"grade,omitempty"`
	StartTime       *time.Time     `json:"start_time,omitempty"`
	EndTime         *time.Time     `json:"end_time,omitempty"`
	IsPublic        bool           `gorm:"default:false" json:"is_public"`
	EngineVersion   string         `json:"engine_version,omitempty"`
	CriteriaVersion string         `json:"criteria_version,omitempty"`
	RawResponse     datatypes.JSON `gorm:"type:jsonb" json:"raw_response,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (a *Analysis) UpdateFromAPIResponse(response interface{}) error {
	rawJSON, err := json.Marshal(response)
	if err != nil {
		return err
	}
	a.RawResponse = datatypes.JSON(rawJSON)
	return nil
}

func (a *Analysis) IsInProgress() bool {
	return a.Status == "DNS" || a.Status == "IN_PROGRESS"
}

func (a *Analysis) IsCompleted() bool {
	return a.Status == "READY" || a.Status == "ERROR"
}

func (a *Analysis) HasError() bool {
	return a.Status == "ERROR"
}

func (a *Analysis) IsReady() bool {
	return a.Status == "READY"
}
