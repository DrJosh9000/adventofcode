package main

import (
	"fmt"
	"os"
	"strings"
)

// Advent of Code 2015
// Day 10

// It's not *that* huge...

func compress(s string) string {
	sb := new(strings.Builder)
	f, n, d := false, 0, rune(0)
	for _, c := range s {
		if !f {
			f, n, d = true, 1, c
			continue
		}
		if c != d {
			sb.WriteRune(rune(n) + '0')
			sb.WriteRune(d)
			n, d = 1, c
			continue
		}
		n++
	}
	sb.WriteRune(rune(n) + '0')
	sb.WriteRune(d)
	return sb.String()
}

func main() {
	s := os.Args[1]
	for i := 0; i < 40; i++ {
		s = compress(s)
	}
	fmt.Println(len(s))
	for i := 0; i < 10; i++ {
		s = compress(s)
	}
	fmt.Println(len(s))
}
