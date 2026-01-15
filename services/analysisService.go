package services

import (
	"fmt"
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"
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
		return nil, fmt.Errorf("failed to start SSL Labs analysis: %w", err)
	}

	analysis := s.createAnalysisFromResponse(resp)

	now := time.Now()
	analysis.LastCheckedAt = &now

	if err := s.repo.Create(analysis); err != nil {
		return nil, fmt.Errorf("failed to save analysis: %w", err)
	}

	if !analysis.IsCompleted() {
		s.PollAnalysisInBackground(analysis.ID)
	}

	return analysis, nil
}

func (s *AnalysisService) UpdateAnalysisStatus(analysisID uuid.UUID) (*models.Analysis, error) {
	analysis, err := s.repo.FindByID(analysisID)
	if err != nil {
		return nil, fmt.Errorf("analysis not found: %w", err)
	}

	if analysis.IsCompleted() {
		return analysis, nil
	}

	request := dto.AnalysisRequest{
		Host: analysis.Host,
		All:  "done",
	}

	resp, err := s.ssl.Analyze(request)
	if err != nil {
		return nil, fmt.Errorf("failed to check SSL Labs status: %w", err)
	}

	s.updateAnalysisFromResponse(analysis, resp)

	now := time.Now()
	analysis.LastCheckedAt = &now

	if err := s.repo.Update(analysis); err != nil {
		return nil, fmt.Errorf("failed to update analysis: %w", err)
	}

	return analysis, nil
}

func (s *AnalysisService) createAnalysisFromResponse(resp *dto.AnalysisResponse) *models.Analysis {
	startTime := time.UnixMilli(resp.StartTime)

	analysis := &models.Analysis{
		ID:              uuid.New(),
		Host:            resp.Host,
		Status:          resp.Status,
		StartTime:       &startTime,
		IsPublic:        resp.IsPublic,
		EngineVersion:   resp.EngineVersion,
		CriteriaVersion: resp.CriteriaVersion,
	}

	return analysis
}

func (s *AnalysisService) updateAnalysisFromResponse(analysis *models.Analysis, resp *dto.AnalysisResponse) {
	analysis.Status = resp.Status

	if resp.TestTime > 0 {
		testTime := time.UnixMilli(resp.TestTime)
		analysis.TestTime = &testTime
	}
}

func (s *AnalysisService) PollAnalysisInBackground(analysisID uuid.UUID) {
	go func() {
		for {
			analysis, err := s.repo.FindByID(analysisID)
			if err != nil {
				return
			}

			if analysis.IsCompleted() {
				return
			}

			delay := constants.GetPollingDelay(analysis.Status)

			time.Sleep(time.Duration(delay) * time.Second)

			request := dto.AnalysisRequest{
				Host: analysis.Host,
				All:  "done",
			}

			resp, err := s.ssl.Analyze(request)
			if err != nil {
				continue
			}

			s.updateAnalysisFromResponse(analysis, resp)

			now := time.Now()
			analysis.LastCheckedAt = &now

			if err := s.repo.Update(analysis); err != nil {
				continue
			}
		}
	}()
}

func (s *AnalysisService) GetAnalysisById(analysisID uuid.UUID) (*models.Analysis, error) {
	analysis, err := s.repo.FindByID(analysisID)
	if err != nil {
		return nil, fmt.Errorf("analysis not found: %w", err)
	}

	return analysis, nil
}
