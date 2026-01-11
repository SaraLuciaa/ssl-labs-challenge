package constants

const (
	PollingDelayDNS        = 5
	PollingDelayInProgress = 10
	PollingDelayCompleted  = 0
)

func GetPollingDelay(status string) int {
	switch status {
	case DNS:
		return PollingDelayDNS
	case InProgress:
		return PollingDelayInProgress
	case Ready, Error:
		return PollingDelayCompleted
	default:
		return PollingDelayInProgress
	}
}
