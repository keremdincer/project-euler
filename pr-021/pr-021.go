package main

import "fmt"

func d(num int) int {
	var sumOfDivisors = 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sumOfDivisors += i
		}
	}
	return sumOfDivisors
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		if d(d(i)) == i && d(i) != i {
			fmt.Println(d(i), " -> ", i)
			sum += i
		}
	}

	fmt.Println(sum)
}
