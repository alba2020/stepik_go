package main

import "fmt"

func main() {
	var x, p, y, n int
	fmt.Scan(&x, &p, &y)

	for x < y {
		x += x * p / 100
		n++
	}

	fmt.Println(n)
}
