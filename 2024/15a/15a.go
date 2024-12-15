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
	G := grid.BytesFromStrings(strings.Split(whmap, "\n"))

	var robot image.Point
	for p, c := range G.All {
		if c == '@' {
			robot = p
			break
		}
	}

	for _, i := range insts {
		d, ok := algo.CGVL[i]
		if !ok {
			continue
		}
		p := robot
	pushLoop:
		for {
			p = p.Add(d)
			switch G.At(p) {
			case '#':
				break pushLoop
			case 'O':
				continue
			case '.':
				for {
					q := p.Sub(d)
					if q == robot {
						G.Set(p, '@')
						G.Set(q, '.')
						robot = p
						break pushLoop
					}
					G.Set(p, G.At(q))
					p = q
				}
			}
		}
	}

	sum := 0
	for p, c := range G.All {
		if c == 'O' {
			sum += 100*p.Y + p.X
		}
	}
	fmt.Println(sum)
}
