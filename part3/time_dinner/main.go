package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const layout = "2006-01-02 15:04:05"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	
	t1, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}

	dinner := time.Date(
		t1.Year(),
		t1.Month(),
		t1.Day(),
		13,
		00,
		00,
		00,
		t1.Location(),
	)

	if t1.After(dinner) {
		fmt.Println(t1.Add(24 * time.Hour).Format(layout))
	} else {
		fmt.Println(t1.Format(layout))
	}
}
