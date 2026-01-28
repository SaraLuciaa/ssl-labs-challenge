package services

import (
	"errors"
	"testing"
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/dto"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSSLLabsService struct {
	mock.Mock
}

func (m *MockSSLLabsService) Analyze(request dto.AnalysisRequest) (*dto.AnalysisResponse, error) {
	args := m.Called(request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.AnalysisResponse), args.Error(1)
}

func (m *MockSSLLabsService) GetLocationById(ip string) ([]string, error) {
	args := m.Called(ip)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}

type MockAnalysisRepository struct {
	mock.Mock
}

func (m *MockAnalysisRepository) Create(analysis *models.Analysis) error {
	args := m.Called(analysis)
	return args.Error(0)
}

func (m *MockAnalysisRepository) Update(analysis *models.Analysis) error {
	args := m.Called(analysis)
	return args.Error(0)
}

func (m *MockAnalysisRepository) FindByID(id uuid.UUID) (*models.Analysis, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Analysis), args.Error(1)
}

func (m *MockAnalysisRepository) FindAll() ([]models.Analysis, error) {
	args := m.Called()
	return args.Get(0).([]models.Analysis), args.Error(1)
}

type MockEndpointRepository struct {
	mock.Mock
}

func (m *MockEndpointRepository) Create(endpoint *models.Endpoint) error {
	args := m.Called(endpoint)
	return args.Error(0)
}

func (m *MockEndpointRepository) Update(endpoint *models.Endpoint) error {
	args := m.Called(endpoint)
	return args.Error(0)
}

func (m *MockEndpointRepository) FindByAnalysisIDAndIP(analysisID uuid.UUID, ipAddress string) (*models.Endpoint, error) {
	args := m.Called(analysisID, ipAddress)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Endpoint), args.Error(1)
}

func (m *MockEndpointRepository) DeleteByAnalysisID(analysisID uuid.UUID) error {
	args := m.Called(analysisID)
	return args.Error(0)
}

func TestGetAllAnalyses_Success(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	expectedAnalyses := []models.Analysis{
		{
			ID:     uuid.New(),
			Host:   "example.com",
			Status: "READY",
		},
		{
			ID:     uuid.New(),
			Host:   "test.com",
			Status: "IN_PROGRESS",
		},
	}

	mockAnalysisRepo.On("FindAll").Return(expectedAnalyses, nil)

	analyses, err := service.GetAllAnalyses()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(analyses))
	assert.Equal(t, expectedAnalyses, analyses)
	mockAnalysisRepo.AssertExpectations(t)
}

func TestGetAllAnalyses_Error(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	mockAnalysisRepo.On("FindAll").Return([]models.Analysis{}, errors.New("database error"))

	analyses, err := service.GetAllAnalyses()

	assert.Error(t, err)
	assert.Nil(t, analyses)
	assert.Contains(t, err.Error(), "failed to retrieve analyses")
	mockAnalysisRepo.AssertExpectations(t)
}

func TestStartAnalysis_Success(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	request := dto.AnalysisRequest{
		Host:     "example.com",
		StartNew: "on",
		All:      "done",
	}

	sslResponse := &dto.AnalysisResponse{
		Host:            "example.com",
		Port:            443,
		Protocol:        "http",
		IsPublic:        true,
		Status:          "READY",
		EngineVersion:   "2.4.1",
		CriteriaVersion: "2009q",
		StartTime:       time.Now().UnixMilli(),
		TestTime:        time.Now().UnixMilli(),
		Endpoints:       []dto.EndpointDto{},
	}

	mockSSL.On("Analyze", request).Return(sslResponse, nil)
	mockAnalysisRepo.On("Create", mock.AnythingOfType("*models.Analysis")).Return(nil)

	analysis, err := service.StartAnalysis(request)

	assert.NoError(t, err)
	assert.NotNil(t, analysis)
	assert.Equal(t, "example.com", analysis.Host)
	assert.Equal(t, "READY", analysis.Status)
	assert.NotNil(t, analysis.LastCheckedAt)
	mockSSL.AssertExpectations(t)
	mockAnalysisRepo.AssertExpectations(t)
}

func TestStartAnalysis_SSLLabsError(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	request := dto.AnalysisRequest{
		Host: "example.com",
	}

	mockSSL.On("Analyze", request).Return(nil, errors.New("SSL Labs API error"))

	analysis, err := service.StartAnalysis(request)

	assert.Error(t, err)
	assert.Nil(t, analysis)
	assert.Contains(t, err.Error(), "failed to start SSL Labs analysis")
	mockSSL.AssertExpectations(t)
}

func TestStartAnalysis_CreateError(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	request := dto.AnalysisRequest{
		Host: "example.com",
	}

	sslResponse := &dto.AnalysisResponse{
		Host:   "example.com",
		Status: "READY",
	}

	mockSSL.On("Analyze", request).Return(sslResponse, nil)
	mockAnalysisRepo.On("Create", mock.AnythingOfType("*models.Analysis")).Return(errors.New("database error"))

	analysis, err := service.StartAnalysis(request)

	assert.Error(t, err)
	assert.Nil(t, analysis)
	assert.Contains(t, err.Error(), "failed to save analysis")
	mockSSL.AssertExpectations(t)
	mockAnalysisRepo.AssertExpectations(t)
}

func TestGetAnalysisById_Success(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	analysisID := uuid.New()
	expectedAnalysis := &models.Analysis{
		ID:     analysisID,
		Host:   "example.com",
		Status: "READY",
	}

	mockAnalysisRepo.On("FindByID", analysisID).Return(expectedAnalysis, nil)

	analysis, err := service.GetAnalysisById(analysisID)

	assert.NoError(t, err)
	assert.NotNil(t, analysis)
	assert.Equal(t, analysisID, analysis.ID)
	assert.Equal(t, "example.com", analysis.Host)
	mockAnalysisRepo.AssertExpectations(t)
}

func TestGetAnalysisById_NotFound(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	analysisID := uuid.New()

	mockAnalysisRepo.On("FindByID", analysisID).Return(nil, errors.New("record not found"))

	analysis, err := service.GetAnalysisById(analysisID)

	assert.Error(t, err)
	assert.Nil(t, analysis)
	assert.Contains(t, err.Error(), "analysis not found")
	mockAnalysisRepo.AssertExpectations(t)
}

func TestAnalysisFromResponse_NewAnalysis(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	now := time.Now().UnixMilli()
	response := &dto.AnalysisResponse{
		Host:            "example.com",
		Port:            443,
		Protocol:        "http",
		IsPublic:        true,
		Status:          "READY",
		StartTime:       now,
		TestTime:        now,
		EngineVersion:   "2.2.0",
		CriteriaVersion: "2009q",
		Endpoints:       []dto.EndpointDto{},
	}

	analysis := service.AnalysisFromResponse(nil, response)

	assert.NotNil(t, analysis)
	assert.NotEqual(t, uuid.Nil, analysis.ID)
	assert.Equal(t, "example.com", analysis.Host)
	assert.Equal(t, 443, analysis.Port)
	assert.Equal(t, "http", analysis.Protocol)
	assert.True(t, analysis.IsPublic)
	assert.Equal(t, "READY", analysis.Status)
	assert.Equal(t, "2.2.0", analysis.EngineVersion)
	assert.Equal(t, "2009q", analysis.CriteriaVersion)
	assert.NotNil(t, analysis.StartTime)
	assert.NotNil(t, analysis.TestTime)
}

func TestAnalysisFromResponse_UpdateExisting(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	existingID := uuid.New()
	existingAnalysis := &models.Analysis{
		ID:     existingID,
		Host:   "example.com",
		Status: "IN_PROGRESS",
	}

	response := &dto.AnalysisResponse{
		Host:      "example.com",
		Port:      443,
		Status:    "READY",
		Endpoints: []dto.EndpointDto{},
	}

	analysis := service.AnalysisFromResponse(existingAnalysis, response)

	assert.NotNil(t, analysis)
	assert.Equal(t, existingID, analysis.ID)
	assert.Equal(t, "READY", analysis.Status)
	assert.Equal(t, 443, analysis.Port)
}

func TestEndpointsFromResponse_NewEndpoint(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	analysisID := uuid.New()
	endpointDtos := []dto.EndpointDto{
		{
			IPAddress:     "192.0.2.1",
			ServerName:    "example.com",
			StatusMessage: "Ready",
			Grade:         "A+",
			Progress:      100,
		},
	}

	mockEndpointRepo.On("FindByAnalysisIDAndIP", analysisID, "192.0.2.1").Return(nil, errors.New("not found"))
	mockEndpointRepo.On("Create", mock.AnythingOfType("*models.Endpoint")).Return(nil)

	endpoints := service.EndpointsFromResponse(analysisID, endpointDtos)

	assert.Equal(t, 1, len(endpoints))
	assert.Equal(t, "192.0.2.1", endpoints[0].IPAddress)
	assert.Equal(t, "example.com", endpoints[0].ServerName)
	assert.Equal(t, "A+", endpoints[0].Grade)
	assert.Equal(t, 100, endpoints[0].Progress)
	mockEndpointRepo.AssertExpectations(t)
}

func TestEndpointsFromResponse_UpdateExisting(t *testing.T) {
	mockAnalysisRepo := new(MockAnalysisRepository)
	mockEndpointRepo := new(MockEndpointRepository)
	mockSSL := new(MockSSLLabsService)

	service := NewAnalysisService(mockSSL, mockAnalysisRepo, mockEndpointRepo)

	analysisID := uuid.New()
	endpointID := uuid.New()
	existingEndpoint := &models.Endpoint{
		ID:         endpointID,
		AnalysisID: analysisID,
		IPAddress:  "192.0.2.1",
		Grade:      "B",
		Progress:   50,
	}

	endpointDtos := []dto.EndpointDto{
		{
			IPAddress:     "192.0.2.1",
			ServerName:    "example.com",
			StatusMessage: "Ready",
			Grade:         "A+",
			Progress:      100,
		},
	}

	mockEndpointRepo.On("FindByAnalysisIDAndIP", analysisID, "192.0.2.1").Return(existingEndpoint, nil)
	mockEndpointRepo.On("Update", mock.AnythingOfType("*models.Endpoint")).Return(nil)

	endpoints := service.EndpointsFromResponse(analysisID, endpointDtos)

	assert.Equal(t, 1, len(endpoints))
	assert.Equal(t, endpointID, endpoints[0].ID)
	assert.Equal(t, "A+", endpoints[0].Grade)
	assert.Equal(t, 100, endpoints[0].Progress)
	mockEndpointRepo.AssertExpectations(t)
}
