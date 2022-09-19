package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2016
// Day 9, part a

func main() {
	in := bytes.TrimSpace(exp.Must(os.ReadFile("inputs/9.txt")))
	length := 0
	for i := 0; i < len(in); i++ {
		switch in[i] {
		case '(':
			e := bytes.IndexByte(in[i:], ')')
			var l, n int
			exp.Must(fmt.Sscanf(string(in[i+1:][:e]), "%dx%d", &l, &n))
			length += l * n
			i += e + l
		default:
			length++
		}
	}
	fmt.Println(length)
}
