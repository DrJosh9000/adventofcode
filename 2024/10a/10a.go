package main

import (
	_ "embed"
	"fmt"
	"image"
	"math/bits"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

//go:embed inputs/10.txt
var input string

func main() {
	G := grid.BytesFromStrings(exp.NonEmpty(strings.Split(input, "\n")))
	bounds := G.Bounds()

	sets := make([]grid.Dense[uint64], 0, 4)
	for range 4 {
		sets = append(sets, grid.Make[uint64](G.Size()))
	}

	var q []image.Point
	var b int
	for p, c := range G.All {
		if c == '9' {
			bb := b / 64
			sets[bb][p.Y][p.X] = 1 << (b - 64*bb)
			q = append(q, p)
			b++
		}
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, d := range algo.Neigh4 {
			n := p.Add(d)
			if !n.In(bounds) {
				continue
			}
			if G.At(n) != G.At(p)-1 {
				continue
			}
			for i := range sets {
				sets[i][n.Y][n.X] |= sets[i].At(p)
			}
			q = append(q, n)
		}
	}

	sum := 0
	for p, c := range G.All {
		if c == '0' {
			for i := range sets {
				sum += bits.OnesCount64(sets[i].At(p))
			}
		}
	}

	fmt.Println(sum)
}
