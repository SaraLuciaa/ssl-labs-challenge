package dto

import "gorm.io/datatypes"

type EndpointDto struct {
	IPAddress string `json:"ipAddress"`
	ServerName string `json:"serverName,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
	Grade string `json:"grade,omitempty"`
	Progress int `json:"progress"`
	Details datatypes.JSON `json:"details,omitempty"`
}
