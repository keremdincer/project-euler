package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var grid [][]int
var greatestProduct int = 0

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
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x])-3; y++ {
			product := 1
			for i := 0; i < 4; i++ {
				product *= grid[x][y+i]
			}
			if product > greatestProduct {
				greatestProduct = product
			}
		}
	}
}

func verticalProducts() {
	for x := 0; x < len(grid)-3; x++ {
		for y := 0; y < len(grid[x]); y++ {
			product := 1
			for i := 0; i < 4; i++ {
				product *= grid[x+i][y]
			}

			if product > greatestProduct {
				greatestProduct = product
			}
		}
	}
}

func diagonalLeftProducts() {
	for x := 0; x < len(grid)-3; x++ {
		for y := 0; y < len(grid[x])-3; y++ {
			product := 1
			for i := 0; i < 4; i++ {
				product *= grid[x+i][y+i]
			}

			if product > greatestProduct {
				greatestProduct = product
			}
		}
	}
}

func diagonalRightProducts() {
	for x := 0; x < len(grid)-3; x++ {
		for y := 0; y < len(grid[x])-3; y++ {
			product := 1
			for i := 0; i < 4; i++ {
				product *= grid[x+3-i][y+i]
			}

			if product > greatestProduct {
				greatestProduct = product
			}
		}
	}
}

func main() {
	setup()
	horizontalProducts()
	fmt.Println("Greated Product so far:", greatestProduct)
	verticalProducts()
	fmt.Println("Greated Product so far:", greatestProduct)
	diagonalLeftProducts()
	fmt.Println("Greated Product so far:", greatestProduct)
	diagonalRightProducts()
	fmt.Println("Greated Product so far:", greatestProduct)
}
