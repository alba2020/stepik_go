package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	
	d3 := d % 10
	d /= 10
	d2 := d % 10
	d /= 10
	d1 := d % 10

	if d1 != d2 && d2 != d3 && d1 != d3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}