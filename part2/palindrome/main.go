package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString() string {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSuffix(s, "\n")
}

func Reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	reversed := make([]rune, n)
	for i, r := range runes {
		reversed[n - i - 1] = r
	}
	return string(reversed)
}

func main() {
	text := ReadString()

	if text == Reverse(text) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
