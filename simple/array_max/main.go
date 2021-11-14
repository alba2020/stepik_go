package main

import "fmt"

func main() {
	var n, x, positive int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x > 0 {
			positive++
		}
	}

	fmt.Println(positive)
}
