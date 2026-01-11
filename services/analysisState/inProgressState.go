package analysisState

import (
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
)

type InProgressState struct{}

func (s *InProgressState) Handle(analysis *models.Analysis) error {
	analysis.Status = constants.AnalysisStatusInProgress
	analysis.StatusMessage = "Assessment in progress"
	return nil
}

func (s *InProgressState) Name() string {
	return constants.AnalysisStatusInProgress
}

func (s *InProgressState) CanTransitionTo(targetState string) bool {
	return targetState == constants.AnalysisStatusReady ||
		targetState == constants.AnalysisStatusError
}
