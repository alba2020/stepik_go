package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func IsRight(s *string) bool {
	runes := []rune(*s)
	n := len(runes)
	if ! unicode.IsUpper(runes[0]) {
		return false
	}
	if runes[n-1] != '.' {
		return false
	}
	return true
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if IsRight(&text) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
