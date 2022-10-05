package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2016
// Day 18

func evolve(dst, src []byte) int {
	safe := len(src)
	n1 := len(src) - 1
	dst[0] = src[1]
	if dst[0] == '^' {
		safe--
	}
	dst[n1] = src[n1-1]
	if dst[n1] == '^' {
		safe--
	}
	for i := 1; i < n1; i++ {
		dst[i] = '.'
		if src[i-1] != src[i+1] {
			dst[i] = '^'
			safe--
		}
	}
	return safe
}

func main() {
	l1 := bytes.TrimSpace(exp.Must(os.ReadFile("inputs/18.txt")))
	l2 := make([]byte, len(l1))
	sum := 0
	i := 0
	for ; i < 40; i++ {
		sum += evolve(l2, l1)
		l1, l2 = l2, l1
	}
	fmt.Println(sum)
	for ; i < 400_000; i++ {
		sum += evolve(l2, l1)
		l1, l2 = l2, l1
	}
	fmt.Println(sum)
}
