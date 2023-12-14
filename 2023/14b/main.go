package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2023
// Day 14, part b

const inputPath = "2023/inputs/14.txt"

func main() {
	g := exp.MustReadByteGrid(inputPath)

	loads := make([]int, 1, 101)
	for i := 0; i < 100; i++ {
		g = spinCycle(g)
		loads = append(loads, load(g))
	}

	fmt.Println(algo.CyclicPredict(loads, 1_000_000_000))
}

func load(g grid.Dense[byte]) int {
	sum := 0
	for i, row := range g {
		sum += algo.Count(row, 'O') * (len(g) - i)
	}
	return sum
}

func spinCycle(g grid.Dense[byte]) grid.Dense[byte] {
	for i := 0; i < 4; i++ {
		rollNorth(g)
		g = g.RotateCW()
	}
	return g
}

func rollNorth(g grid.Dense[byte]) {
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
}
