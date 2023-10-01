package enum

import "errors"

type MarketLifecycleState string

const (
	MarketLifecycleStateNone           MarketLifecycleState = "none"
	MarketLifecycleStateRegistration   MarketLifecycleState = "registration"
	MarketLifecycleStateInitialization MarketLifecycleState = "initialization"
	MarketLifecycleStateCompleted      MarketLifecycleState = "completed"
	MarketLifecycleStateReady          MarketLifecycleState = "ready"
	MarketLifecycleStateRunning        MarketLifecycleState = "running"
	MarketLifecycleStateStopped        MarketLifecycleState = "stoped"
	MarketLifecycleStateFinished       MarketLifecycleState = "finished"
)

func (e *MarketLifecycleState) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for MarketLifecycleState enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = MarketLifecycleStateNone
	case "registration":
		*e = MarketLifecycleStateRegistration
	case "initialization":
		*e = MarketLifecycleStateInitialization
	case "completed":
		*e = MarketLifecycleStateCompleted
	case "ready":
		*e = MarketLifecycleStateReady
	case "running":
		*e = MarketLifecycleStateRunning
	case "stoped":
		*e = MarketLifecycleStateStopped
	case "finished":
		*e = MarketLifecycleStateFinished
	default:
		return errors.New("invalid scan value '" + enumValue + "' for MarketLifecycleState enum")
	}

	return nil
}

func (e MarketLifecycleState) String() string {
	return string(e)
}
