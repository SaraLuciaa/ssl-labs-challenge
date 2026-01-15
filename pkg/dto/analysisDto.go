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
	IsPublic bool `json:"isPublic"`
	Status string `json:"status"`
	StartTime int64 `json:"startTime"`
	TestTime int64 `json:"testTime,omitempty"`
	EngineVersion string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
	Endpoints []EndpointDto `json:"endpoints,omitempty"`
}
