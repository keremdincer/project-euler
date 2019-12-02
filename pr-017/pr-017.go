package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

// If all the numbers from 1 to 1000 (one thousand) inclusive were written out
// in words, how many letters would be used?

// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
// forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
// letters. The use of "and" when writing out numbers is in compliance with
// British usage.

var m map[int]string

func initMap() {
	m = make(map[int]string)

	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"
	m[6] = "six"
	m[7] = "seven"
	m[8] = "eight"
	m[9] = "nine"
	m[10] = "ten"
	m[11] = "eleven"
	m[12] = "twelve"
	m[13] = "thirteen"
	m[14] = "fourteen"
	m[15] = "fifteen"
	m[16] = "sixteen"
	m[17] = "seventeen"
	m[18] = "eighteen"
	m[19] = "nineteen"
	m[20] = "twenty"
	m[30] = "thirty"
	m[40] = "forty"
	m[50] = "fifty"
	m[60] = "sixty"
	m[70] = "seventy"
	m[80] = "eighty"
	m[90] = "ninety"
	m[100] = "onehundred"
	m[200] = "twohundred"
	m[300] = "threehundred"
	m[400] = "fourhundred"
	m[500] = "fivehundred"
	m[600] = "sixhundred"
	m[700] = "sevenhundred"
	m[800] = "eighthundred"
	m[900] = "ninehundred"
	m[1000] = "onethousand"
}

func numberToDigits(number int) []int {
	var digits []int
	s := strconv.Itoa(number)
	for _, r := range s {
		c := string(r)
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, n)
	}

	return digits
}

func numberToLetter(number int) string {
	match, ok := m[number]
	if ok {
		return match
	}

	// If map does not contain a value for given number, slice first digit from
	// number convert it to its letter form then recall function with remaining
	// number.

	firstDigit := string(strconv.Itoa(number)[0])
	remaining := strconv.Itoa(number)[1:]

	n, err := strconv.Atoi(firstDigit)
	if err != nil {
		log.Fatal(err)
	}
	n = n * int(math.Pow(10, float64(len(remaining))))
	match, ok = m[n]

	r, err := strconv.Atoi(remaining)
	if err != nil {
		log.Fatal(err)
	}

	if n/100 >= 1 && n < 1000 {
		match += "and"
	}

	return match + numberToLetter(r)
}

func main() {
	initMap()

	letterCount := 0
	for i := 1; i <= 1000; i++ {
		letterCount += len(numberToLetter(i))
		fmt.Printf("%d => %s\n", i, numberToLetter(i))
	}

	fmt.Println(letterCount)
}
