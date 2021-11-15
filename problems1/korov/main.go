package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // 0 < n < 100

	if n >= 11 && n <= 14 {
		fmt.Printf("%d korov\n", n)
	} else if n % 10 == 1 {
		fmt.Printf("%d korova\n", n)
	} else if n % 10 >= 2 && n % 10 <= 4 {
		fmt.Printf("%d korovy\n", n)
	} else {
		fmt.Printf("%d korov\n", n)
	}
}
