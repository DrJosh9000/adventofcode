package main

import (
	"fmt"
	"image"
	"sync"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/para"
)

// Advent of Code 2023
// Day 16, part a

const inputPath = "2023/inputs/16.txt"

var up, down, left, right = image.Pt(0, -1), image.Pt(0, 1), image.Pt(-1, 0), image.Pt(1, 0)

func main() {
	g := exp.MustReadByteGrid(inputPath)
	bounds := g.Bounds()

	type beam struct {
		p, d image.Point
	}
	var µ sync.Mutex
	energised := make(algo.Set[image.Point])
	seen := make(algo.Set[beam])
	q := para.NewQueue(beam{d: right})
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

		switch g[b.p.Y][b.p.X] {
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

	fmt.Println(len(energised))
}
