package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

//go:embed inputs/15.txt
var input string

func main() {
	whmap, insts := exp.MustCut(input, "\n\n")
	G1 := grid.BytesFromStrings(strings.Split(whmap, "\n"))
	h, w := G1.Size()
	w *= 2
	G := grid.Make[byte](h, w)
	var robot image.Point
	for p, c := range G1.All {
		p1, p2 := image.Pt(2*p.X, p.Y), image.Pt(2*p.X+1, p.Y)
		switch c {
		case '#', '.':
			G.Set(p1, c)
			G.Set(p2, c)
		case 'O':
			G.Set(p1, '[')
			G.Set(p2, ']')
		case '@':
			G.Set(p1, '@')
			G.Set(p2, '.')
			robot = p1
		}
	}
	// fmt.Println(G)
	// fmt.Println(robot)

	L, R := image.Pt(-1, 0), image.Pt(1, 0)

	for _, i := range insts {
		d, ok := algo.CGVL[i]
		if !ok {
			continue
		}
		moves := make(algo.Set[image.Point])
		moves.Insert(robot)
		queue := []image.Point{robot}
		move := true
	pushLoop:
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			q := p.Add(d)
			switch G.At(q) {
			case '#':
				move = false
				break pushLoop
			case '.':
				continue
			case '[':
				if d.Y == 0 {
					moves.Insert(q)
					queue = append(queue, q)
				} else {
					moves.Insert(q, q.Add(R))
					queue = append(queue, q, q.Add(R))
				}
			case ']':
				if d.Y == 0 {
					moves.Insert(q)
					queue = append(queue, q)
				} else {
					moves.Insert(q, q.Add(L))
					queue = append(queue, q, q.Add(L))
				}
			}
		}
		if !move {
			continue
		}
		robot = robot.Add(d)

		next := G.Clone()
		for p := range moves {
			next.Set(p, '.')
		}
		for p := range moves {
			next.Set(p.Add(d), G.At(p))
		}
		G = next
	}

	// fmt.Println(G)

	sum := 0
	for p, c := range G.All {
		if c == '[' {
			sum += 100*p.Y + p.X
		}
	}
	fmt.Println(sum)
}
