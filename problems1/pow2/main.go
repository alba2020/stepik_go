package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for t:=1; t <= n; t *= 2 {
		fmt.Print(t, " ")
	}
	fmt.Println()
}
