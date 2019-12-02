package main

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1
// to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?

func divisible(number int) bool {
	for i := 2; i < 20; i++ {
		if number%i != 0 {
			return false
		}
	}

	return true
}

func main() {
	result := 0
	number := 20

	for {
		if divisible(number) {
			result = number
			break
		}

		number += 2
	}

	fmt.Print(result)
}
