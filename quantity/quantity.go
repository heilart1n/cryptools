package quantity

import (
	"github.com/shopspring/decimal"
)

const (
	DecimalFormat int32 = 6
)

type (
	Quantity struct {
		// value is always represented exclusively in the common decimal format.
		// That is, when creating a quantity, any of your values will be converted considering the common decimal.
		value   decimal.Decimal
		decimal int32
	}
)

func New(value decimal.Decimal, decimalNative int32, fromHumanReadable bool) Quantity {
	if fromHumanReadable {
		value = value.Shift(decimalNative)
	}
	return Quantity{
		value:   value,
		decimal: decimalNative,
	}
}
