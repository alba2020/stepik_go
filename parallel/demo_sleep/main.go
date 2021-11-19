package main

import (
	"fmt"
	"time"
)

func main() {
	go myFunc()
	time.Sleep(1 * time.Second)
}

func myFunc() {
	fmt.Println("hello")
}
