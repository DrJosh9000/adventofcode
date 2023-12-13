package main

import (
	"fmt"
	"slices"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
	"github.com/DrJosh9000/exp/para"
)

// Advent of Code 2023
// Day 13, part b

const inputPath = "2023/inputs/13.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	var g []string
	var gs []grid.Dense[byte]
	for _, line := range lines {
		if line == "" {
			gs = append(gs, grid.BytesFromStrings(g))
			g = nil
		} else {
			g = append(g, line)
		}
	}
	gs = append(gs, grid.BytesFromStrings(g))
	fmt.Println(algo.Sum(para.Map(gs, func(g grid.Dense[byte]) int {
		orows := find(g)
		gt := g.Transpose()
		ocols := find(gt)
		row, col := -1, -1
		if len(orows) == 1 {
			row = orows[0]
		}
		if len(ocols) == 1 {
			col = ocols[0]
		}
		if row == -1 && col == -1 {
			panic("no row or col found for original")
		}
		if row != -1 && col != -1 {
			panic("both a row and a col found for original")
		}

		for i := range g {
			for j := range g[i] {
				flip(&g[i][j])
				flip(&gt[j][i])
				nrows := find(g)
				ncols := find(gt)
				for _, r := range nrows {
					if r != row {
						return 100 * r
					}
				}
				for _, c := range ncols {
					if c != col {
						return c
					}
				}
				flip(&g[i][j])
				flip(&gt[j][i])
			}
		}
		panic("that shouldn't happen")
	})))
}

func flip(b *byte) {
	if *b == '.' {
		*b = '#'
	} else {
		*b = '.'
	}
}

func find(g grid.Dense[byte]) []int {
	var rows []int
	for r := 1; r < len(g); r++ {
		mirror := true
		for i := 0; r-i > 0 && r+i < len(g); i++ {
			j, k := r-i-1, r+i
			if !slices.Equal(g[j], g[k]) {
				mirror = false
			}
		}
		if mirror {
			rows = append(rows, r)
		}
	}
	return rows
}
