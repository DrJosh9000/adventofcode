package main

import (
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2022
// Day 14, part a

func main() {
	cave := make(grid.Sparse[byte])
	ran := algo.Range[int]{-1, 1}
	maxy := 0
	for _, line := range exp.MustReadLines("inputs/14.txt") {
		pts := strings.Split(line, " -> ")
		var p image.Point
		exp.Must(fmt.Sscanf(pts[0], "%d,%d", &p.X, &p.Y))
		if p.Y > maxy {
			maxy = p.Y
		}
		for _, pt := range pts[1:] {
			var q image.Point
			exp.Must(fmt.Sscanf(pt, "%d,%d", &q.X, &q.Y))
			if q.Y > maxy {
				maxy = q.Y
			}
			d := q.Sub(p)
			d.X = ran.Clamp(d.X)
			d.Y = ran.Clamp(d.Y)
			for p != q {
				cave[p] = '#'
				p = p.Add(d)
			}
			cave[q] = '#'
		}
	}

	// Sand pouring
	count := 0
pourLoop:
	for {
		s := image.Pt(500, 0)
		for s.Y <= maxy {
			fall := false
			for _, d := range []image.Point{{0, 1}, {-1, 1}, {1, 1}} {
				t := s.Add(d)
				if cave[t] == 0 {
					fall = true
					s = t
					break
				}
			}
			// Still falling?
			if fall {
				continue
			}
			// Stopped
			//fmt.Println("stopped at", s)
			count++
			cave[s] = 'O'
			continue pourLoop
		}
		// c2, offset := cave.ToDense()
		// c2.Map(func(b byte) byte {
		// 	if b == 0 {
		// 		return ' '
		// 	}
		// 	return b
		// })
		// fmt.Println(offset)
		// fmt.Println(c2)

		fmt.Println(count)
		return
	}
}
