package main

import "fmt"

func main() {
	fmt.Println(fibonacci(6))
}

func fibonacci(n int) int {
	f1 := 1
	f2 := 1
	i := 1
	for i < n {
		f1, f2 = f2, f1 + f2
		i++
	}
	return f1
}
