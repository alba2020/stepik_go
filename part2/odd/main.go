package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	var odd []rune

	for i, r := range s {
		if i % 2 == 1 {
			odd = append(odd, r)
		}
	}

	fmt.Println(string(odd))
}
