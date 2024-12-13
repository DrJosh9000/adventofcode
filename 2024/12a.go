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

//go:embed inputs/12.txt
var input string

func main() {
	G := grid.BytesFromStrings(exp.NonEmpty(strings.Split(input, "\n")))
	bounds := G.Bounds()
	u := make(algo.DisjointSets[image.Point])
	for p, c := range G.All {
		u.Find(p)
		for _, d := range []image.Point{{1, 0}, {0, 1}} {
			q := p.Add(d)
			if !q.In(bounds) {
				continue
			}
			if G.At(q) == c {
				u.Union(p, q)
			}
		}
	}

	sum := 0
	for _, s := range u.Sets() {
		perim := 0
		for _, p := range s {
			c := G.At(p)
			for _, d := range algo.Neigh4 {
				q := p.Add(d)
				if !q.In(bounds) || G.At(q) != c {
					perim++
				}
			}
		}
		sum += perim * len(s)
	}
	fmt.Println(sum)
}
