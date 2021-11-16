package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	for _, r := range s {
		fmt.Print((r - '0') * (r - '0'))
	}
}
