package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 14, part a

const inputPath = "2023/inputs/14.txt"

func main() {
	g := exp.MustReadByteGrid(inputPath)

	for i, row := range g {
		if i == 0 {
			continue
		}
		for j, c := range row {
			if c != 'O' {
				continue
			}
			k := i
			for k > 0 && g[k-1][j] == '.' {
				g[k][j] = '.'
				k--
				g[k][j] = 'O'
			}
		}
	}

	sum := 0
	for i, row := range g {
		sum += algo.Count(row, 'O') * (len(g) - i)
	}

	fmt.Println(sum)
}
