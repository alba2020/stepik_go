package main

import (
	"bufio"
	"os"
	"strconv"
)

// strconv, bufio, os, io

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	sum := 0

	for s.Scan() {
		x, _ := strconv.Atoi(s.Text())
		// w.WriteString(s.Text() + "\n")
		sum += x
		w.Flush()
	}

	w.WriteString(strconv.Itoa(sum))
	w.Flush()
}
