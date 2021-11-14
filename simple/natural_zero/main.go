package main

import "fmt"

func main() {
	var x, max, n int
	
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		if x > max {
			max = x
			n = 1
		} else if x == max {
			n++
		}
	}

	fmt.Println(n)
}
