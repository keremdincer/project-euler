package main

import (
	"fmt"
	"math"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
// that the 6th prime is 13. What is the 10 001st prime number?

func isPrime(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	prime := true
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			prime = false
			break
		}
	}

	return prime
}

func main() {
	var lastPrime int
	primeCounter := 0
	counter := 2

	for {
		if isPrime(counter) {
			lastPrime = counter
			primeCounter++
		}
		if primeCounter == 10001 {
			break
		}
		counter++
	}

	fmt.Println(lastPrime)
}
