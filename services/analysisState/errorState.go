package analysisState

import (
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
)

type ErrorState struct{}

func (s *ErrorState) Handle(analysis *models.Analysis) error {
	now := time.Now()
	analysis.Status = constants.AnalysisStatusError
	if analysis.StatusMessage == "" {
		analysis.StatusMessage = "An error occurred during the assessment"
	}
	analysis.EndTime = &now
	return nil
}

func (s *ErrorState) Name() string {
	return constants.AnalysisStatusError
}

func (s *ErrorState) CanTransitionTo(targetState string) bool {
	return false
}
