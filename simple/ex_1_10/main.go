package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	
	for i := c; i <= n; i += c {
		if i % d != 0 {
			fmt.Println(i)
			break
		}
	}
}
