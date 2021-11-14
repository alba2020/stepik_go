package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("I like Go!")
	}

	var s string = "abcпривет"
	_ = s
}
