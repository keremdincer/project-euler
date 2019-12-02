package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

// Reverse given string
func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func isPalindromic(number int) bool {
	s := strconv.Itoa(number)
	firstHalf := s[:len(s)/2]
	secondHalf := reverse(s[len(s)/2:])
	return firstHalf == secondHalf
}

func main() {
	largestPalindrome := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			if isPalindromic(product) && product > largestPalindrome {
				largestPalindrome = product
			}
		}
	}

	fmt.Println(largestPalindrome)
}
