package enum

import "errors"

type StockMarketType string

const (
	StockMarketTypeNone    StockMarketType = "none"
	StockMarketTypeSpot    StockMarketType = "spot"
	StockMarketTypeFutures StockMarketType = "futures"
	StockMarketTypeOptions StockMarketType = "options"
	StockMarketTypeSwaps   StockMarketType = "swaps"
)

func (e *StockMarketType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for StockMarketType enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = StockMarketTypeNone
	case "spot":
		*e = StockMarketTypeSpot
	case "futures":
		*e = StockMarketTypeFutures
	case "options":
		*e = StockMarketTypeOptions
	case "swaps":
		*e = StockMarketTypeSwaps
	default:
		return errors.New("invalid scan value '" + enumValue + "' for StockMarketType enum")
	}

	return nil
}

func (e StockMarketType) String() string {
	return string(e)
}
