package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		var x int

		select {
		case x = <- firstChan:
			output <- x * x
		case x = <- secondChan:
			output <- x * 3
		case <- stopChan:
			return
		}
		}()

	return output
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	stopChan := make(chan struct{})

	// ch1 <- 12

	res := calculator(ch1, ch2, stopChan)
	ch2 <- 1
	fmt.Println(<-res)
}
