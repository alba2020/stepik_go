package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	var j, k int
	for i := 0; i < 3; i++ {
		fmt.Scan(&j, &k)
		workArray[j], workArray[k] = workArray[k], workArray[j]
	}

	for _, x := range workArray {
		fmt.Printf("%d ", x)
	}
}
