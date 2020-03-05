package main

import (
	"fmt"
	"math"
)

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

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
	sum := 0

	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
