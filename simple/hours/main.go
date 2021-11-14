package main

import "fmt"

func main () {
	var d int
	fmt.Scan(&d)

	minutes_total := d * 2
	minutes := minutes_total % 60
	hours := (minutes_total - minutes) / 60

	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}