package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 10, part b

func main() {
	g := grid.Make[byte](6, 40)
	g.Fill(' ')
	x := 1
	c := 0

	draw := func() {
		if x-1 <= c%40 && x+1 >= c%40 {
			g[c/40][c%40] = '#'
		}
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

	for _, r := range g {
		fmt.Println(string(r))
	}
}
