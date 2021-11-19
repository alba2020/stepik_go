package main

import "fmt"

func main() {

	fn := func(x uint) uint {
		var y uint

		i := 1
		for x > 0 {
			digit := x % 10
			x /= 10
			
			if digit % 2 == 0 && digit != 0 {
				y += digit * uint(i)
				i *= 10
			}
		}

		if y == 0 {
			return 100
		}
		return y
	}

	fmt.Println(fn(10))
}
