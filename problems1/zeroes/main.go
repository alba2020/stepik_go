package main

import "fmt"

func main() {
	var n, x, zeroes int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 0 {
			zeroes++
		}
	}

	fmt.Println(zeroes)
}
