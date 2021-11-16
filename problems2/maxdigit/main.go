package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	var max rune

	for _, r := range s {
		if r > max {
			max = r
		}
	}
	fmt.Println(string(max))
}
