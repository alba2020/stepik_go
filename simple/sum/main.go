package main

import "fmt"

func main() {
	var a, b, s int
	fmt.Scan(&a, &b)

	for s = 0; a <= b; a++ {
		s += a
	}
	
	fmt.Println(s)
}