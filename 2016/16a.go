package main

import (
	"fmt"
	"os"
)

// Advent of Code 2016
// Day 16, part a

// Try to be clever? it turns out that part b isn't as huge as it could be.

func dragon(s string) string {
	b := make([]byte, len(s)*2+1)
	copy(b, s)
	b[len(s)] = '0'
	for i, c := range s {
		b[len(b)-i-1] = '0'
		if c == '0' {
			b[len(b)-i-1] = '1'
		}
	}
	return string(b)
}

func reduce(s string) string {
	b := make([]byte, len(s)/2)
	for i := range b {
		switch s[2*i:][:2] {
		case "00", "11":
			b[i] = '1'
		case "01", "10":
			b[i] = '0'
		}
	}
	return string(b)
}

func main() {
	input := os.Args[1]
	const target = 272

	s := input
	for len(s) < target {
		s = dragon(s)
	}

	s = s[:target]

	for len(s)%2 == 0 {
		s = reduce(s)
	}

	fmt.Println(s)
}
