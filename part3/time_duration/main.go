package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	s = strings.ReplaceAll(s, "мин", "m")
	s = strings.ReplaceAll(s, "сек", "s")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, " ", "")

	dur, _ := time.ParseDuration(s)

	const now = 1589570165
	t := time.Unix(now, 0).UTC()
	t = t.Add(dur)
	fmt.Println(t.Format(time.UnixDate))
}
