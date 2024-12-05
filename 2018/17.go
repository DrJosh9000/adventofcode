package main

import (
	"fmt"
	"image"
	"log"
	"math"

	"drjosh.dev/exp"
)

type span struct {
	min, max int
}

func main() {
	yrange := span{math.MaxInt, math.MinInt}
	plot := make(map[image.Point]struct{})
	exp.MustForEachLineIn("inputs/17.txt", func(line string) {
		var r span
		var t int
		var f, s rune
		if _, err := fmt.Sscanf(line, "%c=%d, %c=%d..%d", &f, &t, &s, &r.min, &r.max); err != nil {
			log.Fatalf("Couldn't scan: %v", err)
		}
		xr, yr := r, span{t, t}
		if f == 'x' {
			xr, yr = yr, xr
		}

		for y := yr.min; y <= yr.max; y++ {
			for x := xr.min; x <= xr.max; x++ {
				plot[image.Pt(x, y)] = struct{}{}
			}
		}
		if yr.min < yrange.min {
			yrange.min = yr.min
		}
		if yr.max > yrange.max {
			yrange.max = yr.max
		}
	})

	oldplot := len(plot)
	up, down, left, right := image.Pt(0, -1), image.Pt(0, 1), image.Pt(-1, 0), image.Pt(1, 0)
	seen := make(map[image.Point]struct{})
	q := map[image.Point]struct{}{{500, 0}: {}}
	for len(q) > 0 {
		var p image.Point
		for q0 := range q {
			p = q0
			break
		}
		delete(q, p)
		if p.Y > yrange.max {
			continue
		}
		if p.Y >= yrange.min {
			seen[p] = struct{}{}
		}

		below := p.Add(down)
		if _, filled := plot[below]; !filled {
			// flow downwards
			q[below] = struct{}{}
			continue
		}

		lw, rw, hole := false, false, false // left wall, right wall
		for p := p.Add(left); ; p = p.Add(left) {
			// hit wall to the left?
			if _, filled := plot[p]; filled {
				lw = true
				break
			}
			seen[p] = struct{}{}
			// hole?
			below := p.Add(down)
			if _, filled := plot[below]; !filled {
				// flow downwards
				q[below] = struct{}{}
				hole = true
				break
			}
		}
		for p := p.Add(right); ; p = p.Add(right) {
			if _, filled := plot[p]; filled {
				rw = true
				break
			}
			seen[p] = struct{}{}
			below := p.Add(down)
			if _, filled := plot[below]; !filled {
				q[below] = struct{}{}
				hole = true
				break
			}
		}

		if lw && rw && !hole {
			// go back up one, and fill the line
			q[p.Add(up)] = struct{}{}
			plot[p] = struct{}{}
			for p := p.Add(left); ; p = p.Add(left) {
				if _, filled := plot[p]; filled {
					break
				}
				plot[p] = struct{}{}
			}
			for p := p.Add(right); ; p = p.Add(right) {
				if _, filled := plot[p]; filled {
					break
				}
				plot[p] = struct{}{}
			}
		}
	}
	fmt.Println("Reachable tiles:", len(seen))
	fmt.Println("Retained:", len(plot)-oldplot)
}
