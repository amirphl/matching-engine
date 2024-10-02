package enum

import "errors"

type SelfTradePreventionType string

const (
	SelfTradePreventionTypeNone        SelfTradePreventionType = "none"
	SelfTradePreventionTypeCancelTaker SelfTradePreventionType = "cancel_taker"
	SelfTradePreventionTypeCancelMaker SelfTradePreventionType = "cancel_maker"
	SelfTradePreventionTypeCancelAll   SelfTradePreventionType = "cancel_all"
)

func (e *SelfTradePreventionType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for SelfTradePreventionType enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = SelfTradePreventionTypeNone
	case "cancel_taker":
		*e = SelfTradePreventionTypeCancelTaker
	case "cancel_maker":
		*e = SelfTradePreventionTypeCancelMaker
	case "cancel_all":
		*e = SelfTradePreventionTypeCancelAll
	default:
		return errors.New("invalid scan value '" + enumValue + "' for SelfTradePreventionType enum")
	}

	return nil
}

func (e SelfTradePreventionType) String() string {
	return string(e)
}
