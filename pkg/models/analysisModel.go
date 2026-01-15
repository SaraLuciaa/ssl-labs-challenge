package models

import (
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"
	"github.com/google/uuid"
)

type Analysis struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Host string `gorm:"not null" json:"host"`
	Port string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	IsPublic bool `gorm:"default:false" json:"is_public"`
	Status string `gorm:"not null" json:"status"`
	StartTime *time.Time `json:"start_time,omitempty"`
	TestTime *time.Time `json:"test_time,omitempty"`
	EngineVersion string `json:"engine_version,omitempty"`
	CriteriaVersion string  `json:"criteria_version,omitempty"`
	LastCheckedAt *time.Time `json:"last_checked_at,omitempty"`
	Endpoints []Endpoint `gorm:"foreignKey:AnalysisID;constraint:OnDelete:CASCADE" json:"endpoints,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Analysis) IsInProgress() bool {
	return a.Status == constants.DNS || a.Status == constants.InProgress
}

func (a *Analysis) IsCompleted() bool {
	return a.Status == constants.Ready || a.Status == constants.Error
}

func (a *Analysis) HasError() bool {
	return a.Status == constants.Error
}

func (a *Analysis) IsReady() bool {
	return a.Status == constants.Ready
}
