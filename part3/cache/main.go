package main

import "fmt"

func work(x int) int {
	return x + 1
}

func main() {
	var x int
	cache := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&x)

		if _, ok := cache[x]; !ok {
			cache[x] = work(x)
		}
		fmt.Print(cache[x], " ")
	}
}
