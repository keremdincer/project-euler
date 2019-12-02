package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

// Work out the first ten digits of the sum of the following one-hundred
// 50-digit numbers.

func main() {

	content, err := ioutil.ReadFile("pr-013/digits.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")
	sum := new(big.Int)

	for _, line := range lines {
		bigNumber := new(big.Int)
		bigNumber, ok := bigNumber.SetString(line, 10)
		if ok {
			sum.Add(sum, bigNumber)
		}
	}

	fmt.Println(sum)
}
