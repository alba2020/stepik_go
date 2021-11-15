package main

import "fmt"

func main() {
	var s, x string
	fmt.Scan(&s, &x)
	
	for _, r := range s {
		if string(r) != x {
			fmt.Print(string(r))
		}
	}
	fmt.Println()
}
