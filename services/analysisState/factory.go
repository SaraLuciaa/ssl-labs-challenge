package analysisState

import "github.com/SaraLuciaa/ssl-labs-challenge/pkg/constants"

func FromStatus(status string) AnalysisState {
	switch status {
	case constants.AnalysisStatusDNS:
		return &DNSState{}
	case constants.AnalysisStatusInProgress:
		return &InProgressState{}
	case constants.AnalysisStatusReady:
		return &ReadyState{}
	case constants.AnalysisStatusError:
		return &ErrorState{}
	default:
		return &ErrorState{}
	}
}
