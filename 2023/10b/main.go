package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2023
// Day 10, part b

const inputPath = "2023/inputs/10.txt"

func main() {
	G := exp.MustReadByteGrid(inputPath)
	h, w := G.Size()

	var start image.Point
startLoop:
	for y, row := range G {
		for x, c := range row {
			if c == 'S' {
				start = image.Pt(x, y)
				break startLoop
			}
		}
	}

	bounds := G.Bounds()

	neigh := map[byte][]image.Point{
		'|': {{0, -1}, {0, 1}},
		'-': {{-1, 0}, {1, 0}},
		'L': {{0, -1}, {1, 0}},
		'J': {{0, -1}, {-1, 0}},
		'7': {{0, 1}, {-1, 0}},
		'F': {{0, 1}, {1, 0}},
		'S': algo.Neigh4,
	}

	// corners of cell
	// corners := []image.Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}

	// cells surrounding point
	surrounds := []image.Point{{-1, -1}, {-1, 0}, {0, -1}, {0, 0}}

	xwest := grid.Make[bool](h, w+1)
	xnorth := grid.Make[bool](h+1, w)

	algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
		switch G[p.Y][p.X] {
		case 'F':
			xwest[p.Y][p.X+1] = true
			xnorth[p.Y+1][p.X] = true

		case 'L':
			xwest[p.Y][p.X+1] = true
			xnorth[p.Y][p.X] = true

		case '7':
			xwest[p.Y][p.X] = true
			xnorth[p.Y+1][p.X] = true

		case 'J':
			xwest[p.Y][p.X] = true
			xnorth[p.Y][p.X] = true

		case '-':
			xwest[p.Y][p.X] = true
			xwest[p.Y][p.X+1] = true

		case '|':
			xnorth[p.Y][p.X] = true
			xnorth[p.Y+1][p.X] = true

		case 'S':
			xwest[p.Y][p.X] = true
			// xnorth[p.Y][p.X] = true
			xwest[p.Y][p.X+1] = true
			// xnorth[p.Y+1][p.X] = true
		}

		var next []image.Point
		for _, dt := range neigh[G[p.Y][p.X]] {
			if r := p.Add(dt); r.In(bounds) {
				valid := false
				for _, dr := range neigh[G[r.Y][r.X]] {
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

	cornsIn := grid.Make[int](h, w)
	var q []image.Point
	for x := range xwest {
		q = append(q, image.Pt(x, 0), image.Pt(x, h))
	}
	for y := range xnorth {
		q = append(q, image.Pt(0, y), image.Pt(w, y))
	}

	seen := make(map[image.Point]bool)

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if seen[p] {
			continue
		}
		seen[p] = true

		if p.X < 0 || p.X > w || p.Y < 0 || p.Y > h {
			continue
		}

		for _, ds := range surrounds {
			if r := p.Add(ds); r.In(bounds) {
				cornsIn[r.Y][r.X]++
			}
		}

		if p.Y > 0 && !xwest[p.Y-1][p.X] {
			q = append(q, p.Add(image.Pt(0, -1)))
		}
		if p.X < w && !xnorth[p.Y][p.X] {
			q = append(q, p.Add(image.Pt(1, 0)))
		}
		if p.Y < h && !xwest[p.Y][p.X] {
			q = append(q, p.Add(image.Pt(0, 1)))
		}
		if p.X > 0 && !xnorth[p.Y][p.X-1] {
			q = append(q, p.Add(image.Pt(-1, 0)))
		}
	}

	count := 0
	for _, row := range cornsIn {
		for _, c := range row {
			if c == 0 {
				count++
				// G[i][j] = 'I'
			}
		}
	}

	// fmt.Println(xwest)
	// fmt.Println(xnorth)
	// fmt.Println(G)
	fmt.Println(count)
}
