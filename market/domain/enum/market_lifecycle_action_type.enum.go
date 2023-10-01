package enum

import "errors"

type MarketLifecycleActionType string

const (
	MarketLifecycleActionTypeNone              MarketLifecycleActionType = "none"
	MarketLifecycleActionTypeSubmit            MarketLifecycleActionType = "submit"
	MarketLifecycleActionTypeReinit            MarketLifecycleActionType = "reinit"
	MarketLifecycleActionTypePause             MarketLifecycleActionType = "pause"
	MarketLifecycleActionTypeResume            MarketLifecycleActionType = "resume"
	MarketLifecycleActionTypeStop              MarketLifecycleActionType = "stop"
	MarketLifecycleActionTypeArchive           MarketLifecycleActionType = "archive"
	MarketLifecycleActionTypeSettle            MarketLifecycleActionType = "settle"
	MarketLifecycleActionTypeCancel            MarketLifecycleActionType = "cancel"
	MarketLifecycleActionTypeSchedule          MarketLifecycleActionType = "schedule"
	MarketLifecycleActionTypeSuccessTransition MarketLifecycleActionType = "success_transition"
	MarketLifecycleActionTypeFailureTransition MarketLifecycleActionType = "failure_transition"
)

func (e *MarketLifecycleActionType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for MarketLifecycleActionType enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = MarketLifecycleActionTypeNone
	case "submit":
		*e = MarketLifecycleActionTypeSubmit
	case "reinit":
		*e = MarketLifecycleActionTypeReinit
	case "pause":
		*e = MarketLifecycleActionTypePause
	case "resume":
		*e = MarketLifecycleActionTypeResume
	case "stop":
		*e = MarketLifecycleActionTypeStop
	case "archive":
		*e = MarketLifecycleActionTypeArchive
	case "settle":
		*e = MarketLifecycleActionTypeSettle
	case "cancel":
		*e = MarketLifecycleActionTypeCancel
	case "schedule":
		*e = MarketLifecycleActionTypeSchedule
	case "success_transition":
		*e = MarketLifecycleActionTypeSuccessTransition
	case "failure_transition":
		*e = MarketLifecycleActionTypeFailureTransition
	default:
		return errors.New("invalid scan value '" + enumValue + "' for MarketLifecycleActionType enum")
	}

	return nil
}

func (e MarketLifecycleActionType) String() string {
	return string(e)
}
