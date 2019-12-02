package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var grid [][]int

func setup() {
	content, err := ioutil.ReadFile("pr-011/grid.txt")
	if err != nil {
		log.Fatal(err)
	}

	textContent := string(content)

	lines := strings.Split(textContent, "\n")

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var row []int

		for _, number := range numbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			row = append(row, value)
		}

		grid = append(grid, row)
	}
}

func horizontalProducts() {
	for _, x := range grid {
		for y := 0; y < len(x)-3; y++ {

		}
	}
}

func verticalProducts() {

}

func diagonalRightProducts() {

}

func diagonalLeftProducts() {

}

func main() {
	setup()
	horizontalProducts()
}
