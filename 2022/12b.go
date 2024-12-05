package main

import (
	"errors"
	"fmt"
	"image"
	"math"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2022
// Day 12, part b

func main() {
	g := exp.Must(grid.BytesFromStrings(exp.MustReadLines("inputs/12.txt")))

	bounds := g.Bounds()

	starts := make(algo.Set[image.Point])
	var end image.Point
	for y, row := range g {
		for x, c := range row {
			switch c {
			case 'E':
				end = image.Pt(x, y)
			case 'S':
				g[y][x] = 'a'
				fallthrough
			case 'a':
				starts.Insert(image.Pt(x, y))
			}
		}
	}

	g[end.Y][end.X] = 'z'

	mindist := math.MaxInt

	for start := range starts {
		algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
			if p == end {
				if d < mindist {
					mindist = d
				}
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

	fmt.Println(mindist)
}
