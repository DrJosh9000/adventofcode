package main

import (
	"fmt"
	"image"
	"iter"
	"math"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 16, part a

const inputPath = "2024/inputs/16.txt"

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

	type state struct {
		p, d image.Point
	}

	best := math.MaxInt
	algo.Dijkstra(
		state{p: start, d: image.Pt(-1, 0)},
		func(s state, score int) (iter.Seq2[state, int], error) {
			return func(yield func(state, int) bool) {
				if s.p == end {
					best = min(best, score)
					return
				}
				yield(state{p: s.p, d: image.Pt(s.d.Y, -s.d.X)}, 1000)
				yield(state{p: s.p, d: image.Pt(-s.d.Y, s.d.X)}, 1000)
				if q := s.p.Add(s.d); q.In(bounds) && G.At(q) != '#' {
					yield(state{p: q, d: s.d}, 1)
				}
			}, nil
		},
	)
	fmt.Println(best)
}
