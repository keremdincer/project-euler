package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
)

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
	number := new(big.Int)
	numberTwo := new(big.Int)

	number.SetInt64(2)
	numberTwo.SetInt64(2)

	for i := 0; i < 999; i++ {
		number.Mul(number, numberTwo)
	}
	fmt.Println(digitSum(number.Text(10)))
}
