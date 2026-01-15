package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Endpoint struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	AnalysisID uuid.UUID `gorm:"type:uuid;not null" json:"analysis_id"`
	IPAddress string `json:"ip_address,omitempty"`
	ServerName string `json:"server_name,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Grade string `json:"grade,omitempty"`
	Progress int `json:"progress,omitempty"`
	Details datatypes.JSON `gorm:"type:jsonb" json:"details,omitempty"`
}
