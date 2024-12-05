package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2023
// Day 23, part a

const inputPath = "2023/inputs/23.txt"

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

		var q image.Point
		switch g.At(p) {
		case '>':
			q = p.Add(image.Pt(1, 0))

		case '<':
			q = p.Add(image.Pt(-1, 0))

		case '^':
			q = p.Add(image.Pt(0, -1))

		case 'v':
			q = p.Add(image.Pt(0, 1))

		case '.':
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
			return
		}

		if been.At(q) {
			return
		}
		been[q.Y][q.X] = true
		search(append(path, q))
		been[q.Y][q.X] = false
	}

	search([]image.Point{start})

	fmt.Println(longest)
}
