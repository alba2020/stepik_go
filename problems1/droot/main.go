package main

import "fmt"

func droot(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if sum < 10 {
		return sum
	} else {
		return droot(sum)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(droot(n))
}