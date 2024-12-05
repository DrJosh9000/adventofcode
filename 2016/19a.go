package main

import (
	"fmt"
	"os"
	"strconv"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 19, part a

func main() {
	n := exp.Must(strconv.Atoi(os.Args[1]))

	e := make([]int, n)
	for i := range e[1:] {
		e[i] = i + 1
	}
	e[n-1] = 0

	p := 0
	for ; e[p] != p; p = e[p] {
		e[p] = e[e[p]]
	}

	fmt.Println(p + 1)
}
