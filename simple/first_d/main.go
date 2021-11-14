package main

import "fmt"

func main() {
	var d int // d <= 10_000
	fmt.Scan(&d)
	
	for d >= 10 {
		d -= d % 10
		d /= 10
	}
	
	fmt.Println(d)
}
