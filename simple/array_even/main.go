package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	slice := make([]int, n, n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	
	for i := 0; i < n; i += 2 {
		fmt.Print(slice[i], " ")
	}

	fmt.Println()
}