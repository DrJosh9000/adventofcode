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

//go:embed inputs/10.txt
var input string

func main() {
	G := grid.BytesFromStrings(exp.NonEmpty(strings.Split(input, "\n")))
	bounds := G.Bounds()

	score := grid.Make[int](G.Size())
	var q []image.Point
	for p, c := range G.All {
		if c == '0' {
			score[p.Y][p.X] = 1
			q = append(q, p)
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
			sum := 0
			for _, dd := range algo.Neigh4 {
				nn := n.Add(dd)
				if !nn.In(bounds) {
					continue
				}
				if G.At(nn) == G.At(n)+1 {
					sum += score.At(nn)
				}
			}
			score[n.Y][n.X] = sum
			q = append(q, n)
		}
	}

	// fmt.Println(score)

	sum := 0
	for p, c := range G.All {
		if c == '0' {
			sum += score.At(p)
		}
	}

	fmt.Println(sum)
}
