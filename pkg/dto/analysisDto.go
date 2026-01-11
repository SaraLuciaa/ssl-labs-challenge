package dto

type AnalysisRequest struct {
	Host           string `json:"host" binding:"required"`
	Publish        string `json:"publish"`
	StartNew       string `json:"start_new"`
	FromCache      string `json:"from_cache"`
	MaxAge         int    `json:"max_age"`
	All            string `json:"all"`
	IgnoreMismatch string `json:"ignore_mismatch"`
}

type AnalysisResponse struct {
	Host            string            `json:"host"`
	Port            int               `json:"port"`
	Protocol        string            `json:"protocol"`
	IsPublic        bool              `json:"isPublic"`
	Status          string            `json:"status"`
	StatusMessage   string            `json:"statusMessage"`
	StartTime       int64             `json:"startTime"`
	TestTime        int64             `json:"testTime,omitempty"`
	EngineVersion   string            `json:"engineVersion"`
	CriteriaVersion string            `json:"criteriaVersion"`
	Endpoints       []EndpointSummary `json:"endpoints,omitempty"`
	CertHostnames   []string          `json:"certHostnames,omitempty"`
}

type EndpointSummary struct {
	IPAddress     string `json:"ipAddress"`
	ServerName    string `json:"serverName,omitempty"`
	StatusMessage string `json:"statusMessage"`
	Grade         string `json:"grade,omitempty"`
	HasWarnings   bool   `json:"hasWarnings,omitempty"`
	IsExceptional bool   `json:"isExceptional,omitempty"`
	Progress      int    `json:"progress"`
	Duration      int    `json:"duration,omitempty"`
	Eta           int    `json:"eta,omitempty"`
}
