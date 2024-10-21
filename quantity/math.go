package quantity

//
//func (q Quantity) Mul(q2 Quantity) Quantity {
//	return Quantity{
//		value:         q.value.Mul(q2.value).Shift(-DecimalCommon),
//		decimalNative: q.decimalNative,
//		decimalCommon: q.decimalCommon,
//	}
//}
//
//func (q Quantity) Div(q2 Quantity) Quantity {
//	return Quantity{
//		value:         q.value.Div(q2.value).Shift(DecimalCommon),
//		decimalNative: q.decimalNative,
//		decimalCommon: q.decimalCommon,
//	}
//}
//
//func (q Quantity) Sub(q2 Quantity) Quantity {
//	return Quantity{
//		value:         q.value.Sub(q2.value),
//		decimalNative: q.decimalNative,
//		decimalCommon: q.decimalCommon,
//	}
//}
//
//func (q Quantity) Add(q2 Quantity) Quantity {
//	return Quantity{
//		value:         q.value.Add(q2.value),
//		decimalNative: q.decimalNative,
//		decimalCommon: q.decimalCommon,
//	}
//}
//
//func (q Quantity) PercentageOf(total Quantity) decimal.Decimal {
//	if total.value.IsZero() {
//		return decimal.Zero
//	}
//	percentage := q.value.Div(total.value).Mul(decimal.NewFromInt(100))
//	return percentage
//}
//
//// GreaterThan проверяет, является ли текущее значение больше другого
//// Пример:
//// v1 := NewValue(200)
//// v2 := NewValue(100)
//// result := v1.GreaterThan(v2) // result будет true
//func (q Quantity) GreaterThan(q2 Quantity) bool {
//	return q.value.GreaterThan(q2.value)
//}
//
//// LessThan проверяет, является ли текущее значение меньше другого
//// Пример:
//// v1 := NewValue(100)
//// v2 := NewValue(200)
//// result := v1.LessThan(v2) // result будет true
//func (q Quantity) LessThan(q2 Quantity) bool {
//	return q.value.LessThan(q2.value)
//}
//
//// Equal проверяет, равняется ли текущее значение другому
//// Пример:
//// v1 := NewValue(100)
//// v2 := NewValue(100)
//// result := v1.Equal(v2) // result будет true
//func (q Quantity) Equal(q2 Quantity) bool {
//	return q.value.Equal(q2.value)
//}
//
//// NotEqual проверяет, не равняется ли текущее значение другому
//// Пример:
//// v1 := NewValue(100)
//// v2 := NewValue(200)
//// result := v1.NotEqual(v2) // result будет true
//func (q Quantity) NotEqual(q2 Quantity) bool {
//	return !q.value.Equal(q2.value)
//}
//
//// GreaterThanOrEqual проверяет, является ли текущее значение больше или равно другому
//// Пример:
//// v1 := NewValue(200)
//// v2 := NewValue(200)
//// result := v1.GreaterThanOrEqual(v2) // result будет true
//func (q Quantity) GreaterThanOrEqual(q2 Quantity) bool {
//	return q.value.GreaterThanOrEqual(q2.value)
//}
//
//// LessThanOrEqual проверяет, является ли текущее значение меньше или равно другому
//// Пример:
//// v1 := NewValue(100)
//// v2 := NewValue(200)
//// result := v1.LessThanOrEqual(v2) // result будет true
//func (q Quantity) LessThanOrEqual(q2 Quantity) bool {
//	return q.value.LessThanOrEqual(q2.value)
//}
