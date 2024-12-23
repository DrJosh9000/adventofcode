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
// Day 16, part b

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
	bestscore := math.MaxInt
	var q []state

	prev := exp.Must(algo.Dijkstra(
		state{p: start, d: image.Pt(-1, 0)},
		func(s state, score int) (iter.Seq2[state, int], error) {
			return func(yield func(state, int) bool) {
				if s.p == end {
					switch {
					case score < bestscore:
						bestscore = score
						q = []state{s}
					case score == bestscore:
						q = append(q, s)
					}
					return
				}
				yield(state{p: s.p, d: image.Pt(s.d.Y, -s.d.X)}, 1000)
				yield(state{p: s.p, d: image.Pt(-s.d.Y, s.d.X)}, 1000)
				if q := s.p.Add(s.d); q.In(bounds) && G.At(q) != '#' {
					yield(state{p: q, d: s.d}, 1)
				}
			}, nil
		},
	))

	best := make(algo.Set[image.Point])
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		best.Insert(s.p)
		q = append(q, prev[s]...)
	}

	// for b := range best {
	// 	G.Set(b, 'O')
	// }
	// fmt.Println(G)

	fmt.Println(len(best))
}
