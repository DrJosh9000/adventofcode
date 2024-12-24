package main

import (
	"fmt"
	"image"
	"iter"
	"math"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2024
// Day 18, part a

const inputPath = "2024/inputs/18.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	obs := make(algo.Set[image.Point])
	for _, line := range lines[:1024] {
		var p image.Point
		fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
		obs.Insert(p)
	}

	bounds := image.Rect(0, 0, 71, 71)
	best := math.MaxInt
	end := image.Pt(70, 70)
	pred, _ := algo.AStar(
		image.Pt(0, 0),
		func(p image.Point) int { return algo.L1(p.Sub(end)) },
		func(p image.Point, d int) (iter.Seq2[image.Point, int], error) {
			if p == end {
				best = min(best, d)
				return nil, nil
			}
			return func(yield func(image.Point, int) bool) {
				for _, e := range algo.Neigh4 {
					q := p.Add(e)
					if q.In(bounds) && !obs.Contains(q) {
						if !yield(q, 1) {
							return
						}
					}
				}
			}, nil
		},
	)

	G := grid.Make[byte](71, 71)
	G.Fill('.')
	for p := range obs {
		G.Set(p, '#')
	}
	for p := end; p != image.Pt(0, 0); p = pred[p][0] {
		G.Set(p, 'O')
	}
	fmt.Println(G)

	fmt.Println(best)
}
