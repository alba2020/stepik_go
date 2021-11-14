package main

import "fmt"

func main() {
	var n, sum, x int
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x > 9 && x < 100 && x % 8 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
