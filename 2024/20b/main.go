package main

import (
	"fmt"
	"image"
	"iter"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2024
// Day 20, part b

const inputPath = "2024/inputs/20.txt"

func main() {
	G := exp.MustReadByteGrid(inputPath)
	bounds := G.Bounds()
	var start, end image.Point
	for p, c := range G.All {
		switch c {
		case 'S':
			start = p
		case 'E':
			end = p
		}
	}

	D := grid.Make[int](G.Size())
	prev, _ := algo.FloodFill(start, func(p image.Point, d int) (iter.Seq[image.Point], error) {
		return func(yield func(p image.Point) bool) {
			D.Set(p, d)
			for _, d := range algo.Neigh4 {
				if q := p.Add(d); q.In(bounds) && G.At(q) != '#' {
					if !yield(q) {
						return
					}
				}
			}
		}, nil
	})

	cheats := make(algo.Set[cheat])
	var path []image.Point
	for p := end; ; p = prev[p][0] {
		path = append(path, p)
		if p == start {
			break
		}
	}

	for p, q := range algo.AllPairs(path) {
		d := p.Sub(q)
		m := algo.L1(d)
		if m > 20 {
			continue
		}
		if algo.Abs(D.At(p)-D.At(q)) >= 100+m {
			cheats.Insert(order(p, q))
		}
	}

	fmt.Println(len(cheats))
}

type cheat struct{ s, e image.Point }

func order(p, q image.Point) cheat {
	if p.X < q.X || (p.X == q.X && p.Y < q.Y) {
		return cheat{p, q}
	}
	return cheat{q, p}
}
