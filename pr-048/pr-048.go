package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := new(big.Int)

	for i := 1; i <= 1000; i++ {
		num := new(big.Int)
		pow := new(big.Int)
		pow.SetString("1", 10)
		num.SetInt64(int64(i))

		for j := 1; j <= i; j++ {
			pow.Mul(pow, num)
		}
		sum.Add(sum, pow)
	}
	text := sum.Text(10)
	fmt.Println(text[len(text)-10:])
}
