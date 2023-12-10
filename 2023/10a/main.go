package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 10, part a

const inputPath = "2023/inputs/10.txt"

func main() {
	grid := exp.MustReadByteGrid(inputPath)

	var start image.Point
startLoop:
	for y, row := range grid {
		for x, c := range row {
			if c == 'S' {
				start = image.Pt(x, y)
				break startLoop
			}
		}
	}

	bounds := grid.Bounds()

	neigh := map[byte][]image.Point{
		'|': {{0, -1}, {0, 1}},
		'-': {{-1, 0}, {1, 0}},
		'L': {{0, -1}, {1, 0}},
		'J': {{0, -1}, {-1, 0}},
		'7': {{0, 1}, {-1, 0}},
		'F': {{0, 1}, {1, 0}},
		'S': algo.Neigh4,
	}

	furthest := 0

	algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
		furthest = max(furthest, d)

		var next []image.Point
		for _, dt := range neigh[grid[p.Y][p.X]] {
			if r := p.Add(dt); r.In(bounds) {
				valid := false
				for _, dr := range neigh[grid[r.Y][r.X]] {
					if s := r.Add(dr); s == p {
						valid = true
						break
					}
				}
				if valid {
					next = append(next, r)
				}
			}
		}
		return next, nil
	})
	fmt.Println(furthest)
}
