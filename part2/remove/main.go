package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	var res []rune

	for _, r := range s {
		if strings.Count(s, string(r)) == 1 {
			res = append(res, r)
		}
	}
	
	fmt.Println(string(res))
}
