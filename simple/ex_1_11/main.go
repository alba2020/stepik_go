package main

import "fmt"

func main() {
	outer: for {
		var x int
		fmt.Scan(&x)
		
		switch {
		case x < 10:
			continue
		case x > 100:
			break outer
		default:
			fmt.Println(x)
		}
	}
}
