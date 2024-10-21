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
	q1 := quantity.New(decimal.NewFromFloat(10), 9, true)
	q2 := quantity.New(decimal.NewFromFloat(155), 2, true)
	q3 := quantity.New(decimal.NewFromUint64(963175561288580), 6, false)
	fmt.Printf("Q1: %s\n", q1.StringFormat(2))
	fmt.Printf("Q2: %s\n", q2.StringFormat(2))
	fmt.Printf("Q3: %s\n", q3.StringFormat(6))
}
