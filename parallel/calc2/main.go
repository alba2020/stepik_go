package main

import (
	"fmt"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		var x, sum int

		for {
			select {
			case x = <- arguments:
				sum += x
			case <- done:
				out <- sum
				return
			}
		}
	}()

	return out
}


func main() {
	args := make(chan int)
	done := make(chan struct{})

	res := calculator(args, done)
	
	args <- 1
	args <- 2
	args <- 3

	done <- * new (struct{})

	fmt.Println(<-res)
}
