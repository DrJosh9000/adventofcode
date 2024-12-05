package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 13, part a

const inputPath = "2023/inputs/13.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	var g []string
	var cols, rows int
	for _, line := range lines {
		if line == "" {
			cs, rs := process(g)
			cols += cs
			rows += rs
			g = g[:0]
		} else {
			g = append(g, line)
		}
	}
	cs, rs := process(g)
	cols += cs
	rows += rs
	fmt.Println(cols + 100*rows)
}

func process(g []string) (cols, rows int) {
	rows = find(g)
	g = transpose(g)
	cols = find(g)
	return cols, rows
}

func transpose(g []string) []string {
	h := make([][]byte, len(g[0]))
	for i := range h {
		for j := range g {
			h[i] = append(h[i], g[j][i])
		}
	}
	return algo.Map(h, func(b []byte) string { return string(b) })
}

func find(g []string) int {
	var rows int
	for r := 1; r < len(g); r++ {
		mirror := true
		for i := 0; r-i > 0 && r+i < len(g); i++ {
			j, k := r-i-1, r+i
			if g[j] != g[k] {
				mirror = false
			}
		}
		if mirror {
			rows = r
		}
	}
	return rows
}
