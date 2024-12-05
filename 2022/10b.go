package main

import (
	"fmt"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2022
// Day 10, part b

func main() {
	g := grid.Make[bool](6, 40)
	x := 1
	c := 0

	draw := func() {
		g[c/40][c%40] = x-1 <= c%40 && x+1 >= c%40
	}

	for _, line := range exp.MustReadLines("inputs/10.txt") {
		fs := strings.Fields(line)
		switch fs[0] {
		case "addx":
			for d := 0; d < 2; d++ {
				draw()
				c++
			}
			x += exp.Must(strconv.Atoi(fs[1]))
		case "noop":
			draw()
			c++
		}
	}

	fmt.Println(g)
}
