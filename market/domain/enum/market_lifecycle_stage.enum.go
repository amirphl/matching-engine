package enum

import "errors"

type MarketLifecycleStage string

const (
	MarketLifecycleStageNone        MarketLifecycleStage = "none"
	MarketLifecycleStageEntrance    MarketLifecycleStage = "entrance"
	MarketLifecycleStageRetention   MarketLifecycleStage = "retention"
	MarketLifecycleStageTermination MarketLifecycleStage = "termination"
)

func (e *MarketLifecycleStage) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for MarketLifecycleStage enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = MarketLifecycleStageNone
	case "entrance":
		*e = MarketLifecycleStageEntrance
	case "retention":
		*e = MarketLifecycleStageRetention
	case "termination":
		*e = MarketLifecycleStageTermination
	default:
		return errors.New("invalid scan value '" + enumValue + "' for MarketLifecycleStage enum")
	}

	return nil
}

func (e MarketLifecycleStage) String() string {
	return string(e)
}
