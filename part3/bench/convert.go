package main

import (
	"fmt"
	"strconv"
)

func ConvertFormat(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func ConvertString(i int) string {
	return fmt.Sprintf("%d", i)
}

func main() {
	fmt.Println(ConvertFormat(12))
	fmt.Println(ConvertString(12))

	res1:=strconv.FormatBool(10 == int16(float64(100/10)))
	fmt.Println(res1)

	// res2 := (strconv.FormatBool(true)) == (10 == int(float64(100/10)))
	// fmt.Prinlnt(res2)
}
