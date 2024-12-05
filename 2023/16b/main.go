package main

import (
	"fmt"
	"image"
	"sync"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
	"drjosh.dev/exp/para"
)

// Advent of Code 2023
// Day 16, part b

const inputPath = "2023/inputs/16.txt"

var up, down, left, right = image.Pt(0, -1), image.Pt(0, 1), image.Pt(-1, 0), image.Pt(1, 0)

type beam struct {
	p, d image.Point
}

func main() {
	g := exp.MustReadByteGrid(inputPath)
	h, w := g.Size()

	best := 0
	for x := range g[0] {
		best = max(best,
			countEnergised(g, beam{p: image.Pt(x, 0), d: down}),
			countEnergised(g, beam{p: image.Pt(x, h-1), d: up}),
		)
	}
	for y := range g {
		best = max(best,
			countEnergised(g, beam{p: image.Pt(0, y), d: right}),
			countEnergised(g, beam{p: image.Pt(w-1, y), d: left}),
		)
	}
	fmt.Println(best)
}

func countEnergised(g grid.Dense[byte], b beam) int {
	bounds := g.Bounds()

	var µ sync.Mutex
	energised := make(algo.Set[image.Point])
	seen := make(algo.Set[beam])
	q := para.NewQueue(b)
	q.Process(func(b beam) {
		if !b.p.In(bounds) {
			return
		}

		µ.Lock()
		if seen.Contains(b) {
			µ.Unlock()
			return
		}
		seen.Insert(b)
		energised.Insert(b.p)
		µ.Unlock()

		switch g.At(b.p) {
		case '.':
			b.p = b.p.Add(b.d)
			q.Push(b)

		case '\\':
			switch b.d {
			case up:
				b.d = left
			case down:
				b.d = right
			case right:
				b.d = down
			case left:
				b.d = up
			}
			b.p = b.p.Add(b.d)
			q.Push(b)

		case '/':
			switch b.d {
			case up:
				b.d = right
			case down:
				b.d = left
			case right:
				b.d = up
			case left:
				b.d = down
			}
			b.p = b.p.Add(b.d)
			q.Push(b)

		case '|':
			switch b.d {
			case left, right:
				q.Push(beam{p: b.p.Add(up), d: up})
				q.Push(beam{p: b.p.Add(down), d: down})

			case up, down:
				b.p = b.p.Add(b.d)
				q.Push(b)
			}

		case '-':
			switch b.d {
			case up, down:
				q.Push(beam{p: b.p.Add(left), d: left})
				q.Push(beam{p: b.p.Add(right), d: right})

			case left, right:
				b.p = b.p.Add(b.d)
				q.Push(b)
			}
		}
	})
	return len(energised)
}
