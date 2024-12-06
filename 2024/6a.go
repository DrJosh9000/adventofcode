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

	dir := image.Pt(0, -1)
	seen := make(algo.Set[image.Point])
	p := start
	for p.In(bounds) {
		seen.Insert(p)

		if q := p.Add(dir); q.In(bounds) && G.At(q) == '#' {
			dir = image.Pt(-dir.Y, dir.X)
		} else {
			p = p.Add(dir)
		}
	}
	fmt.Println(len(seen))
}
