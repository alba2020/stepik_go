package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ones := n % 10
	n = (n - ones) / 10
	tens := n % 10
	
	fmt.Println(tens)
}
