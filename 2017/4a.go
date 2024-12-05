package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

func main() {
	count := 0
	exp.MustForEachLineIn("inputs/4.txt", func(line string) {
		s := make(map[string]bool)
		for _, w := range strings.Fields(line) {
			if s[w] {
				return
			}
			s[w] = true
		}
		count++
	})
	fmt.Println(count)
}