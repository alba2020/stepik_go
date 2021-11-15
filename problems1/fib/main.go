package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	f1 := 1
	f2 := 1
	n := 1

	for {
		if a == f1 {
			fmt.Println(n)
			break
		} else if a < f1 {
			fmt.Println(-1)
			break
		} else {
			f1, f2 = f2, f1 + f2
			n++
		}
	}
}
