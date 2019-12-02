package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number <= 3 {
		return true
	}

	if number%2 == 0 || number%3 == 0 {
		return false
	}

	for i := 5; i*i < number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func largestPrimeFactor(number int) int {
	var largestPrime int
	for i := 2; i < number/2; i++ {
		if number%i == 0 && isPrime(i) {
			largestPrime = i
		}
	}

	return largestPrime
}

// TODO: Refactor functions for efficiency
func main() {
	result := largestPrimeFactor(600851475143)
	fmt.Println(result)
}
