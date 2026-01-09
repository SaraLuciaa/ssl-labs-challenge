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
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	IsPublic        bool   `json:"isPublic"`
	Status          string `json:"status"`
	StatusMessage   string `json:"statusMessage"`
	StartTime       int64  `json:"startTime"`
	EngineVersion   string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
}