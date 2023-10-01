package enum

import "errors"

type CommissionType string

const (
	CommissionTypeNone   CommissionType = "none"
	CommissionTypeFiat   CommissionType = "fiat"
	CommissionTypeCrypto CommissionType = "crypto"
)

func (e *CommissionType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("invalid scan value for CommissionAssetType enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "none":
		*e = CommissionTypeNone
	case "fiat":
		*e = CommissionTypeFiat
	case "crypto":
		*e = CommissionTypeCrypto
	default:
		return errors.New("invalid scan value '" + enumValue + "' for CommissionAssetType enum")
	}

	return nil
}

func (e CommissionType) String() string {
	return string(e)
}
