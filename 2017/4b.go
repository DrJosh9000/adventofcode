package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

type fingerprint [26]int

func badHash(s string) fingerprint {
	var h fingerprint
	for _, c := range s {
		h[c-'a']++
	}
	return h
}

func main() {
	count := 0
	exp.MustForEachLineIn("inputs/4.txt", func(line string) {
		s := make(map[fingerprint]bool)
		for _, w := range strings.Fields(line) {
			h := badHash(w)
			if s[h] {
				return
			}
			s[h] = true
		}
		count++
	})
	fmt.Println(count)
}