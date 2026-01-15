package dto

type AnalysisRequest struct {
	Host string `json:"host" binding:"required"`
	StartNew string `json:"start_new"`
	All string `json:"all"`
}

type AnalysisResponse struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Protocol string `json:"protocol"`
	IsPublic bool `json:"is_public"`
	Status string `json:"status"`
	StartTime int64 `json:"start_time"`
	TestTime int64 `json:"test_time,omitempty"`
	EngineVersion string `json:"engine_version"`
	CriteriaVersion string `json:"criteria_version"`
	Endpoints []EndpointSummary `json:"endpoints,omitempty"`
}

type EndpointSummary struct {
	IPAddress string `json:"ip_address"`
	ServerName string `json:"server_name,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Grade string `json:"grade,omitempty"`
	Progress int `json:"progress"`
	Details map[string]interface{} `json:"details,omitempty"`
}
