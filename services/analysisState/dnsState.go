package analysisState

import (
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
)

type DNSState struct{}

func (s *DNSState) Handle(analysis *models.Analysis) error {
	analysis.Status = constants.AnalysisStatusDNS
	analysis.StatusMessage = "Resolving DNS"
	return nil
}

func (s *DNSState) Name() string {
	return constants.AnalysisStatusDNS
}

func (s *DNSState) CanTransitionTo(targetState string) bool {
	return targetState == constants.AnalysisStatusInProgress ||
		targetState == constants.AnalysisStatusError
}
