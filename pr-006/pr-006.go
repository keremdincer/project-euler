package main

import "fmt"

// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385

// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025

// Hence the difference between the sum of the squares of the first ten natural
// numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.

func sumOfSquares(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i * i
	}
	return sum
}

func squareSumOf(number int) int {
	squareSum := 0
	for i := 1; i <= number; i++ {
		squareSum += i
	}
	return squareSum * squareSum
}

func main() {
	fmt.Printf("%d %d = %d", sumOfSquares(100), squareSumOf(100), (squareSumOf(100) - sumOfSquares(100)))
}
