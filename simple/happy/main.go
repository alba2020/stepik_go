package main

import "fmt"

func main() {
	var n int // d <= 10_000
	fmt.Scan(&n)
	
	f := n % 10; n /= 10
	e := n % 10; n /= 10
	d := n % 10; n /= 10
	c := n % 10; n /= 10
	b := n % 10; n /= 10
	a := n % 10;

	if a + b + c == d + e + f {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
