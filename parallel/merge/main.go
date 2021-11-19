package main

import (
	"fmt"
	"sync"
	"time"
)

func work(fn func(int) int, arg int, res *int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := fn(arg)
	// fmt.Println("got result", num)
	*res = num
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int) {
	go func() {
		wg := new (sync.WaitGroup)
		res1 := make([]int, 10)
		res2 := make([]int, 10)

		for i := 0; i < n; i++ {
			wg.Add(1)
			go work(fn, <-in1, &res1[i], wg)
			wg.Add(1)
			go work(fn, <-in2, &res2[i], wg)
		}
		wg.Wait()

		for i := 0; i < n; i++ {
			sum := res1[i] + res2[i]
			// fmt.Println("writing sum", sum)
			out <- sum
		}
	}()
}

func square(x int) int {
	time.Sleep(1 * time.Second)
	return x * x
}

func main() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)

	merge2Channels(square, in1, in2, out, 3)

	in1 <- 1
	in2 <- 1

	in1 <- 2
	in2 <- 2

	in1 <- 3
	in2 <- 3

	for i := 0; i < 3; i++ {
		// fmt.Println("reading out", i)
		fmt.Println(<-out)
	}
}
