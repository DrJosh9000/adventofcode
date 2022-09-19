package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2016
// Day 9, part b

func length(data []byte) int {
	ln := 0
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '(':
			e := bytes.IndexByte(data[i:], ')')
			var l, n int
			exp.Must(fmt.Sscanf(string(data[i+1:][:e]), "%dx%d", &l, &n))
			ln += length(data[i+e+1:][:l]) * n
			i += e + l
		default:
			ln++
		}
	}
	return ln
}

func main() {
	fmt.Println(length(bytes.TrimSpace(exp.Must(os.ReadFile("inputs/9.txt")))))
}
