package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 0))
}

/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */

 func sumInt(xs... int) (int, int) {
    sum := 0
    for _, el := range xs {
        sum += el
    }
    return len(xs), sum
}
