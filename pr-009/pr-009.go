package main

import (
	"fmt"
	"math"
)

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func main() {
	for c := 2; c < 1000; c++ {
		for b := 1; b < c; b++ {
			for a := 0; a < b; a++ {
				isPythagoreanTriplet := (math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) == math.Pow(float64(c), 2)
				isValid := (a + b + c) == 1000

				if isPythagoreanTriplet && isValid {
					fmt.Println("a:", a, "b:", b, "c:", c)
					fmt.Println(a * b * c)
				}
			}
		}
	}
}
