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
	ssl *SslLabsService
	analysisRepo repositories.AnalysisRepository
	endpointRepo repositories.EndpointRepository
}

func NewAnalysisService(ssl *SslLabsService, analysisRepo repositories.AnalysisRepository, endpointRepo repositories.EndpointRepository) *AnalysisService {
	return &AnalysisService{
		ssl: ssl,
		analysisRepo: analysisRepo,
		endpointRepo: endpointRepo,
	}
}

func (s *AnalysisService) StartAnalysis(request dto.AnalysisRequest) (*models.Analysis, error) {
	resp, err := s.ssl.Analyze(request)
	if err != nil {
		return nil, fmt.Errorf("failed to start SSL Labs analysis: %w", err)
	}

	analysis := s.AnalysisFromResponse(nil, resp)

	now := time.Now()
	analysis.LastCheckedAt = &now

	if err := s.analysisRepo.Create(analysis); err != nil {
		return nil, fmt.Errorf("failed to save analysis: %w", err)
	}

	if !analysis.IsCompleted() {
		s.PollAnalysisInBackground(analysis.ID)
	}

	return analysis, nil
}

func (s *AnalysisService) PollAnalysisInBackground(analysisID uuid.UUID) {
	go func() {
		for {
			analysis, err := s.analysisRepo.FindByID(analysisID)
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

			analysis = s.AnalysisFromResponse(analysis, resp)

			now := time.Now()
			analysis.LastCheckedAt = &now

			if err := s.analysisRepo.Update(analysis); err != nil {
				continue
			}
		}
	}()
}

func (s *AnalysisService) AnalysisFromResponse(analysis *models.Analysis, resp *dto.AnalysisResponse) *models.Analysis {
	if analysis == nil {
		analysis = &models.Analysis{
			ID: uuid.New(),
		}
	}

	analysis.Host = resp.Host
	analysis.Port = resp.Port
	analysis.Protocol = resp.Protocol
	analysis.IsPublic = resp.IsPublic
	analysis.Status = resp.Status
	analysis.EngineVersion = resp.EngineVersion
	analysis.CriteriaVersion = resp.CriteriaVersion

	if resp.StartTime > 0 {
		startTime := time.UnixMilli(resp.StartTime)
		analysis.StartTime = &startTime
	}

	if resp.TestTime > 0 {
		testTime := time.UnixMilli(resp.TestTime)
		analysis.TestTime = &testTime
	}

	analysis.Endpoints = s.EndpointsFromResponse(analysis.ID, resp.Endpoints)

	return analysis
}

func (s *AnalysisService) EndpointsFromResponse(analysisID uuid.UUID, endpointsDto []dto.EndpointDto) []models.Endpoint {
	endpoints := []models.Endpoint{}

	for _, ep := range endpointsDto {
		existingEndpoint, err := s.endpointRepo.FindByAnalysisIDAndIP(analysisID, ep.IPAddress)
		isNewEndpoint := err != nil || existingEndpoint == nil

		var endpoint models.Endpoint
		if err == nil && existingEndpoint != nil {
			endpoint = *existingEndpoint
			isNewEndpoint = false
		} else {
			endpoint = models.Endpoint{
				ID:         uuid.New(),
				AnalysisID: analysisID,
			}
			isNewEndpoint = true
		}

		endpoint.IPAddress = ep.IPAddress
		endpoint.ServerName = ep.ServerName
		endpoint.StatusMessage = ep.StatusMessage
		endpoint.Grade = ep.Grade
		endpoint.Progress = ep.Progress
		endpoint.Details = ep.Details

		if isNewEndpoint {
			s.endpointRepo.Create(&endpoint)
		} else {
			s.endpointRepo.Update(&endpoint)
		}

		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}

func (s *AnalysisService) GetAnalysisById(analysisID uuid.UUID) (*models.Analysis, error) {
	analysis, err := s.analysisRepo.FindByID(analysisID)
	if err != nil {
		return nil, fmt.Errorf("analysis not found: %w", err)
	}

	return analysis, nil
}