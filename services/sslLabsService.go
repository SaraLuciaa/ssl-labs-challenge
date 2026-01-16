package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/dto"
)

type SSLLabsAPI interface {
	Analyze(request dto.AnalysisRequest) (*dto.AnalysisResponse, error)
}

type SslLabsService struct {
	client *http.Client
}

func NewSSLLabsService(client *http.Client) *SslLabsService {
	return &SslLabsService{
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (s *SslLabsService) Analyze(request dto.AnalysisRequest) (*dto.AnalysisResponse, error) {
	url := "https://api.ssllabs.com/api/v2/analyze?host=" + request.Host

	if request.StartNew != "" {
		url += "&startNew=" + request.StartNew
	}
	if request.All != "" {
		url += "&all=" + request.All
	}

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.AnalysisResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
