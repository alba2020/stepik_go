package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const wrong = "Wrong password"
const ok = "Ok"

func main() {
	var password string
	fmt.Scan(&password)

	if utf8.RuneCountInString(password) < 5 {
		fmt.Println(wrong)
		return
	}

	for _, r := range password {
		if unicode.IsDigit(r) {
			continue
		} else if unicode.IsLetter(r) && unicode.Is(unicode.Latin, r) {
			continue
		} else {
			fmt.Println(wrong)
			return
		}
	}
	
	fmt.Println(ok)
}
