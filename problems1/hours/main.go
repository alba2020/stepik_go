package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	h_part := k - (k % 3600)
	h := h_part / 3600

	m_part := k - h_part
	m := (m_part - m_part % 60) / 60

	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
