package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadString() string {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	return strings.TrimSuffix(s, "\n")
}

func main() {
	s := ReadString()

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	
	two := strings.Split(s, ";")

	first, _ := strconv.ParseFloat(two[0], 64)
	second, _ := strconv.ParseFloat(two[1], 64)

	fmt.Printf("%.4f\n", first / second)
}
