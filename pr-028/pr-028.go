package main

import "fmt"

func main() {

	incr := 2
	sum := 0
	counter := 0

	for i := 1; i <= 1001*1001; i += incr {
		sum += i

		if counter != 0 && counter%4 == 0 {
			incr += 2
		}

		counter++
	}

	fmt.Println(sum)

}
