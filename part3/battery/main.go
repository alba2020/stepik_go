package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	charge int
}

func (b Battery) String() string {
	return fmt.Sprintf("[%s%s]",
		strings.Repeat(" ", 10 - b.charge), strings.Repeat("X", b.charge))
}

func main() {
	var input string
	fmt.Scan(&input)

	charge := 0
	for _, r := range input {
		if r == '1' {
			charge++
		}
	}
	batteryForTest := Battery{charge}
	fmt.Println(batteryForTest)
}