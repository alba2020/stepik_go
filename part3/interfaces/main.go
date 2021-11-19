package main

import "fmt"

func readTask() (interface{}, interface{}, interface{}) {
	return 34.1, 2.12, "o"
}

func main() {
	val1, val2, op := readTask()

	const unknown = "неизвестная операция"
	var a, b, res float64

	if v, ok := val1.(float64); ok {
		a = v
	} else {
		fmt.Printf("value=%v: %T\n", val1, val1)
		return
	}

	if v, ok := val2.(float64); ok {
		b = v
	} else {
		fmt.Printf("value=%v: %T\n", val2, val2)
		return
	}

	if v, ok := op.(string); ok {
		switch v {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		default:
			fmt.Printf(unknown)
			return
		}
	} else {
		fmt.Printf(unknown)
		return
	}

	fmt.Printf("%.4f\n", res)
}
