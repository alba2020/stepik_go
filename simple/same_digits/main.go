package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	
	skipx := true
	for divx := 10000; divx > 1; divx /= 10 {
		digitx := (x % divx) / (divx / 10)
		if (digitx != 0) {
			skipx = false
		}
		if (! skipx) {
	
			skipy := true
			for divy := 10000; divy > 1; divy /= 10 {
				digity := (y % divy) / (divy / 10)
				if (digity != 0) {
					skipy = false
				}
				if (! skipy) {
					if digitx == digity {
						fmt.Print(digitx, " ")
						break
					}
				}
			}

		}
	}
	fmt.Println()
}
