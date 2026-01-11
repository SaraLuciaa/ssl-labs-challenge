package analysisState

import "github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"

type AnalysisState interface {
	Handle(analysis *models.Analysis) error

	Name() string

	CanTransitionTo(targetState string) bool
}
