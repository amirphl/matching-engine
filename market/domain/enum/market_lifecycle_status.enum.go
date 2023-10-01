package enum

import "errors"

type MarketLifecycleStatus string

const (
	MarketLifecycleStatusNone         MarketLifecycleStatus = "none"
	MarketLifecycleStatusSuccess      MarketLifecycleStatus = "success"
	MarketLifecycleStatusAccepted     MarketLifecycleStatus = "accepted" //
	MarketLifecycleStatusFailed       MarketLifecycleStatus = "failed"
	MarketLifecycleStatusPending      MarketLifecycleStatus = "pending"
	MarketLifecycleStatusInProgress   MarketLifecycleStatus = "in-progress"
	MarketLifecycleStatusActivated    MarketLifecycleStatus = "activated"
	MarketLifecycleStatusStopped      MarketLifecycleStatus = "stoped"
	MarketLifecycleStatusExpired      MarketLifecycleStatus = "expired"
	MarketLifecycleStatusSuspend      MarketLifecycleStatus = "suspend"
	MarketLifecycleStatusSchedule     MarketLifecycleStatus = "schedule"
	MarketLifecycleStatusRejected     MarketLifecycleStatus = "rejected"
	MarketLifecycleStatusReconfigured MarketLifecycleStatus = "reconfigure" //
	MarketLifecycleStatusCanceled     MarketLifecycleStatus = "canceled"
	MarketLifecycleStatusSettled      MarketLifecycleStatus = "settled"
	MarketLifecycleStatusArchived     MarketLifecycleStatus = "archived"
)

func (e *MarketLifecycleStatus) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for MarketLifecycleStatus enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = MarketLifecycleStatusNone
	case "success":
		*e = MarketLifecycleStatusSuccess
	case "accepted":
		*e = MarketLifecycleStatusAccepted
	case "failed":
		*e = MarketLifecycleStatusFailed
	case "pending":
		*e = MarketLifecycleStatusPending
	case "in-progress":
		*e = MarketLifecycleStatusInProgress
	case "activated":
		*e = MarketLifecycleStatusActivated
	case "stoped":
		*e = MarketLifecycleStatusStopped
	case "expired":
		*e = MarketLifecycleStatusExpired
	case "suspend":
		*e = MarketLifecycleStatusSuspend
	case "schedule":
		*e = MarketLifecycleStatusSchedule
	case "rejected":
		*e = MarketLifecycleStatusRejected
	case "reconfigure":
		*e = MarketLifecycleStatusReconfigured
	case "canceled":
		*e = MarketLifecycleStatusCanceled
	case "settled":
		*e = MarketLifecycleStatusSettled
	case "archived":
		*e = MarketLifecycleStatusArchived
	default:
		return errors.New("invalid scan value '" + enumValue + "' for MarketLifecycleStatus enum")
	}

	return nil
}

func (e MarketLifecycleStatus) String() string {
	return string(e)
}
