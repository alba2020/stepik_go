package main

import (
	"fmt"
	"time"
)

//1986-04-16T05:20:00+06:00
//Wed Apr 16 05:20:00 +0600 1986

func main() {
	var s string
	fmt.Scan(&s)
	
	// layout1 := "2006-01-02T15:04:05-07:00"
	t1, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}

	// layout2 := "Mon Jan 2 15:04:05 -0700 2006"
	fmt.Println(t1.Format(time.UnixDate))
}
