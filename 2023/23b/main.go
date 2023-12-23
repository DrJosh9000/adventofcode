package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2023
// Day 23, part b

const inputPath = "2023/inputs/23.txt"

// Warning:
// "${solution}"  482.28s user 60.83s system 124% cpu 7:15.39 total

func main() {
	g := exp.MustReadByteGrid(inputPath)
	bounds := g.Bounds()
	h, w := g.Size()

	var start, goal image.Point
	for x, c := range g[0] {
		if c == '.' {
			start.X = x
		}
	}
	goal.Y = h - 1
	for x, c := range g[goal.Y] {
		if c == '.' {
			goal.X = x
		}
	}

	been := grid.Make[bool](h, w)
	longest := 0
	var search func([]image.Point)
	search = func(path []image.Point) {
		n1 := len(path) - 1
		p := path[n1]
		if p == goal {
			longest = max(longest, n1)
			// fmt.Println(been)
			return
		}

		for _, d := range algo.Neigh4 {
			q := p.Add(d)
			if !q.In(bounds) {
				continue
			}
			if g.At(q) == '#' {
				continue
			}
			if been.At(q) {
				continue
			}
			been[q.Y][q.X] = true
			search(append(path, q))
			been[q.Y][q.X] = false
		}

	}

	search([]image.Point{start})

	fmt.Println(longest)
}
