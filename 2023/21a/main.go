package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 21, part a

const inputPath = "2023/inputs/21.txt"

func main() {
	g := exp.MustReadByteGrid(inputPath)
	bounds := g.Bounds()
	var start image.Point
startLoop:
	for y, row := range g {
		for x, c := range row {
			if c == 'S' {
				start = image.Pt(x, y)
				g[y][x] = '.'
				break startLoop
			}
		}
	}

	even := make(algo.Set[image.Point])
	algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
		if d > 64 {
			return nil, nil
		}
		if d%2 == 0 {
			even.Insert(p)
		}
		var next []image.Point
		for _, t := range algo.Neigh4 {
			q := p.Add(t)
			if !q.In(bounds) {
				continue
			}
			if g.At(q) == '#' {
				continue
			}
			next = append(next, q)
		}
		return next, nil
	})

	fmt.Println(len(even))
}
