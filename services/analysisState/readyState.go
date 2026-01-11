package analysisState

import (
	"time"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
)

type ReadyState struct{}

func (s *ReadyState) Handle(analysis *models.Analysis) error {
	now := time.Now()
	analysis.Status = constants.AnalysisStatusReady
	analysis.StatusMessage = "Assessment complete"
	analysis.EndTime = &now

	return nil
}

func (s *ReadyState) Name() string {
	return constants.AnalysisStatusReady
}

func (s *ReadyState) CanTransitionTo(targetState string) bool {
	return false
}
