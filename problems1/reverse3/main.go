package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)

	for x > 0 {
		y *= 10
		y += x % 10
		x /= 10
	}

	fmt.Println(y)
}
