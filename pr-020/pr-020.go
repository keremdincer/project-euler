package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
)

// n! means n × (n − 1) × ... × 3 × 2 × 1

// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27

// Find the sum of the digits in the number 100!

func digitSum(numberString string) int {
	number, err := strconv.Atoi(string(numberString[:1]))
	if err != nil {
		log.Fatal(err)
	}
	if len(numberString) == 1 {
		return number
	}
	remaining := numberString[1:]
	return number + digitSum(remaining)
}

func main() {
	product := new(big.Int)
	product.SetString("1", 10)
	for i := 100; i > 0; i-- {
		bigI := new(big.Int)
		bigI.SetInt64(int64(i))
		product.Mul(product, bigI)
	}

	fmt.Println(digitSum(product.Text(10)))
}
