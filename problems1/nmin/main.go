package main

import "fmt"

func main() {
	var n, x, min int
	fmt.Scan(&n)
	
	fmt.Scan(&min)
	count := 1

	for i := 0; i < n - 1; i++ {
		fmt.Scan(&x)
		if x < min {
			min = x
			count = 1
		} else if x == min {
			count++
		}
	}

	fmt.Println(count)
}
