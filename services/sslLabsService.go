package services

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/dto"
)

type SSLLabsAPI interface {
	Analyze(request dto.AnalysisRequest) (*dto.AnalysisResponse, error)
	GetLocationById(ip string) ([]string, error)
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

func (s *SslLabsService) GetLocationById(ip string) ([]string, error) {
	urlll := "http://ip-api.com/csv/" + ip

	response, err := s.client.Get(urlll)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	csvResponse, err := csv.NewReader(response.Body).Read()
	if err != nil {
		return nil, err
	}

	return csvResponse, nil
}

// func parseCSV(data []byte) (*csv.Reader, error) {
//     reader := csv.NewReader(bytes.NewReader(data))
//     return reader, nil
// }
