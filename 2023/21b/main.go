package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 21, part b

const inputPath = "2023/inputs/21.txt"

func main() {
	g := exp.MustReadByteGrid(inputPath)
	h, w := g.Size()
	if h != w {
		panic("input isn't square")
	}
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

	// fmt.Println("grid size", w, "x", h)
	// fmt.Println("start", start)

	// because the map is an odd size, the parity swaps after w...
	// so we have to traverse 2w to cover both parities

	lim := w/2 + 4*w
	dist := make([]int, 3) // w/2, 2*w + w/2, 4*w + w/2.
	algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
		if d > lim {
			return nil, nil
		}
		if d%2 == 1 {
			for i := range dist {
				if want := 2*i*w + w/2; d <= want {
					dist[i]++
				}
			}
		}

		var next []image.Point
		for _, t := range algo.Neigh4 {
			q := p.Add(t)
			if g.At(q.Mod(bounds)) == '#' {
				continue
			}
			next = append(next, q)
		}
		return next, nil
	})

	// Fit ye quadratic
	c := dist[0]
	ba := dist[1] - c
	b2a4 := dist[2] - c
	a := b2a4/2 - ba
	b := ba - a
	fmt.Println(c, b, a)

	x := 26501365 / (2 * w)
	fmt.Println(c + b*x + a*x*x)
}
