package services

import (
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/dto"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
	"github.com/SaraLuciaa/ssl-labs-challenge/repositories"
	"github.com/google/uuid"
)

type AnalysisService struct {
	ssl  *SslLabsService
	repo repositories.AnalysisRepository
}

func NewAnalysisService(ssl *SslLabsService, repo repositories.AnalysisRepository) *AnalysisService {
	return &AnalysisService{
		ssl:  ssl,
		repo: repo,
	}
}

func (s *AnalysisService) StartAnalysis(request dto.AnalysisRequest) (*models.Analysis, error) {
	resp, err := s.ssl.Analyze(request)
	if err != nil {
		return nil, err
	}

	startTime := time.UnixMilli(resp.StartTime)

	analysis := &models.Analysis{
		ID:              uuid.New(),
		Host:            resp.Host,
		Status:          resp.Status,
		StatusMessage:   resp.StatusMessage,
		StartTime:       &startTime,
		IsPublic:        resp.IsPublic,
		EngineVersion:   resp.EngineVersion,
		CriteriaVersion: resp.CriteriaVersion,
	}

	if err := s.repo.Create(analysis); err != nil {
		return nil, err
	}

	return analysis, nil
}
