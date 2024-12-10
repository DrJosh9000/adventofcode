package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

//go:embed inputs/8.txt
var input string

func main() {
	G := grid.BytesFromStrings(strings.Split(input, "\n")[:50])
	bounds := G.Bounds()
	antennae := make(map[byte][]image.Point)
	for p, c := range G.All {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			antennae[c] = append(antennae[c], p)
		}
	}

	antinodes := make(algo.Set[image.Point])
	for _, as := range antennae {
		for a, b := range algo.AllPairs(as) {
			d := a.Sub(b)
			p, q := a.Add(d), b.Sub(d)
			if p.In(bounds) {
				antinodes.Insert(p)
			}
			if q.In(bounds) {
				antinodes.Insert(q)
			}
		}
	}
	fmt.Println(len(antinodes))
}
