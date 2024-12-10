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
			if chr != 'X' {
				continue
			}
			p := image.Pt(x, y)
		dirLoop:
			for _, delta := range algo.Neigh8 {
				q := p
				str := make([]byte, 0, 3)
				for range 3 {
					q = q.Add(delta)
					if !q.In(bounds) {
						continue dirLoop
					}
					str = append(str, G.At(q))
				}
				if string(str) == "MAS" {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
