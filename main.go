package main

import (
	"fmt"
	"github.com/heilart1n/cryptools/quantity"
	_ "github.com/k0kubun/pp"
	"github.com/shopspring/decimal"
)

func main() {
	//q1 := quantity.NewValueIncludeDecimal(0.123456789, 9)
	//fmt.Println(q1.String())
	q1 := quantity.FromHumanReadable(decimal.NewFromFloat(0.1), 9)
	fmt.Println(q1.String())
}
