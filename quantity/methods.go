package quantity

import (
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
	"math/big"
)

func (q Quantity) ToNative() decimal.Decimal {
	return q.value.Shift(q.decimal)
}

func (q Quantity) Float64() float64 {
	return q.ToNative().InexactFloat64()
}

func (q Quantity) Uint64() uint64 {
	return q.ToNative().BigInt().Uint64()
}

func (q Quantity) BigInt() *big.Int {
	return q.ToNative().BigInt()
}

func (q Quantity) Int64() int64 {
	return q.ToNative().BigInt().Int64()
}

func (q Quantity) String() string {
	return q.ToNative().String()
}

func (q Quantity) StringFormat(precision int) string {
	shift := decimal.NewFromInt(10).Pow(decimal.NewFromInt32(q.decimal))
	val := q.value.Div(shift)
	ac := accounting.Accounting{
		Symbol:    "",
		Precision: precision,
		Thousand:  ",",
		Decimal:   ".",
	}
	return ac.FormatMoneyDecimal(val)
}

func (q Quantity) UpdateNativeDecimal(d int32) Quantity {
	q.decimal = d
	return q
}

func (q Quantity) IsZero() bool {
	return q.value.IsZero()
}

func ZeroQuantity() Quantity {
	return Quantity{
		value:   decimal.NewFromInt(0),
		decimal: 2,
	}
}
