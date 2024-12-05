package main

import (
	"errors"
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2022
// Day 12, part a

func main() {
	g := exp.Must(grid.BytesFromStrings(exp.MustReadLines("inputs/12.txt")))

	bounds := g.Bounds()

	var start, end image.Point
	for y, row := range g {
		for x, c := range row {
			switch c {
			case 'S':
				start = image.Pt(x, y)
			case 'E':
				end = image.Pt(x, y)
			}
		}
	}

	g[start.Y][start.X] = 'a'
	g[end.Y][end.X] = 'z'

	algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
		if p == end {
			fmt.Println(d)
			return nil, errors.New("all done")
		}
		var next []image.Point
		for _, step := range algo.Neigh4 {
			t := p.Add(step)
			if !t.In(bounds) {
				continue
			}
			if g[t.Y][t.X] <= g[p.Y][p.X]+1 {
				next = append(next, t)
			}
		}

		return next, nil
	})
}
