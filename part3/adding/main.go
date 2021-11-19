package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(adding("hell123", "sdfs456sljlk"))
}

func OnlyDigits(s string) string {
	res := []rune{}

	for _, r := range s {
		if unicode.IsDigit(r) {
			res = append(res, r)
		}
	}
	return string(res)
}

func adding(a, b string) int64 {
	num1, _ := strconv.Atoi(OnlyDigits(a))
	num2, _ := strconv.Atoi(OnlyDigits(b))
	return int64(num1 + num2)
}
