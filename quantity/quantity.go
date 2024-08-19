package quantity

import (
	"github.com/shopspring/decimal"
)

const (
	DecimalCommon int32 = 18
	DecimalFormat int32 = 6
)

type (
	Quantity struct {
		// value is always represented exclusively in the common decimal format.
		// That is, when creating a quantity, any of your values will be converted considering the common decimal.
		value decimal.Decimal
		// decimalNative is the original decimal of the number, used for displaying native values.
		decimalNative int32
		// decimalCommon is the general decimal used throughout the entire project.
		decimalCommon int32
	}
)

func New(value decimal.Decimal, decimalNative int32, fromHumanReadable bool) Quantity {
	if fromHumanReadable {
		value = value.Shift(DecimalCommon)
	}
	return Quantity{
		value:         value,
		decimalNative: decimalNative,
		decimalCommon: DecimalCommon,
	}
}
