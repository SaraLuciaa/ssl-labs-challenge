package analysisState

import (
	"errors"

	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
)

type AnalysisContext struct {
	currentState AnalysisState
	analysis     *models.Analysis
}

func NewAnalysisContext(analysis *models.Analysis) *AnalysisContext {
	state := FromStatus(analysis.Status)
	return &AnalysisContext{
		currentState: state,
		analysis:     analysis,
	}
}

func (c *AnalysisContext) GetCurrentState() AnalysisState {
	return c.currentState
}

func (c *AnalysisContext) TransitionTo(newStateName string) error {
	if !c.currentState.CanTransitionTo(newStateName) {
		return errors.New("invalid transition from " + c.currentState.Name() + " to " + newStateName)
	}

	c.currentState = FromStatus(newStateName)
	c.analysis.Status = newStateName
	return nil
}

func (c *AnalysisContext) Process() error {
	return c.currentState.Handle(c.analysis)
}

func (c *AnalysisContext) IsTerminal() bool {
	stateName := c.currentState.Name()
	return stateName == "READY" || stateName == "ERROR"
}
