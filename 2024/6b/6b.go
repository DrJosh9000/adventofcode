package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

//go:embed inputs/6.txt
var input string

func main() {
	G := grid.BytesFromStrings(strings.Split(input, "\n"))
	bounds := G.Bounds()
	var start image.Point
startFindLoop:
	for y, row := range G {
		for x, ch := range row {
			if ch == '^' {
				start = image.Pt(x, y)
				break startFindLoop
			}
		}
	}

	obs := make(algo.Set[image.Point])
	for y, row := range G {
		for x, ch := range row {
			if ch == '#' {
				continue
			}
			o := image.Pt(x, y)
			type pair struct {
				p, d image.Point
			}
			path := make(algo.Set[pair])
			p := start
			d := image.Pt(0, -1)
			for p.In(bounds) {
				path.Insert(pair{p, d})
				if q := p.Add(d); q.In(bounds) && G.At(q) == '#' || q == o {
					d = image.Pt(-d.Y, d.X)
				} else {
					p = p.Add(d)
				}
				if path.Contains(pair{p, d}) {
					obs.Insert(o)
					break
				}
			}
		}
	}

	fmt.Println(len(obs))
}
