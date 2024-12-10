package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

//go:embed inputs/4.txt
var input string

func main() {
	G := grid.BytesFromStrings(strings.Split(input, "\n"))
	bounds := G.Bounds()
	count := 0
	for y, row := range G {
		for x, chr := range row {
			if chr != 'A' {
				continue
			}
			p := image.Pt(x, y)
		dirLoop:
			for _, dir := range algo.Neigh4 {
				b0, b1 := p.Sub(dir), p.Add(dir)
				r := image.Pt(-dir.Y, dir.X)
				for _, q := range []image.Point{b0.Add(r), b0.Sub(r), b1.Add(r), b1.Sub(r)} {
					if !q.In(bounds) {
						break dirLoop
					}
				}
				if G.At(b0.Add(r)) == 'M' && G.At(b0.Sub(r)) == 'M' && G.At(b1.Add(r)) == 'S' && G.At(b1.Sub(r)) == 'S' {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
