package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2015
// Day 25

func main() {
	row, col := exp.Must(strconv.Atoi(os.Args[1]))-1, exp.Must(strconv.Atoi(os.Args[2]))-1

	n := uint(col)
	row += col
	for row > 0 {
		n += uint(row)
		row--
	}

	const mod = 33554393
	x := algo.Pow(252533, n, func(x, y int) int {
		return (x * y) % mod
	})

	fmt.Println((x * 20151125) % mod)
}
