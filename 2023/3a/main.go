package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"drjosh.dev/exp"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2023
// Day 3, part a

func main() {
	digitsRE := regexp.MustCompile(`\d+`)

	sum := 0
	g := grid.BytesFromStrings(exp.MustReadLines("2023/inputs/3.txt"))
	bounds := g.Bounds()
	for y, row := range g {
		for _, rg := range digitsRE.FindAllIndex(row, -1) {
			sym := byte('.')

		pointLoop:
			for p := image.Pt(rg[0], y); p.X < rg[1]; p.X++ {
				ds := []image.Point{{0, -1}, {0, 1}}
				if p.X == rg[0] {
					ds = append(ds, image.Pt(-1, -1), image.Pt(-1, 0), image.Pt(-1, 1))
				}
				if p.X == rg[1]-1 {
					ds = append(ds, image.Pt(1, -1), image.Pt(1, 0), image.Pt(1, 1))
				}
				for _, d := range ds {
					q := p.Add(d)
					if !q.In(bounds) {
						continue
					}
					sym = g[q.Y][q.X]
					if sym != '.' {
						break pointLoop
					}
				}
			}

			if sym == '.' {
				continue
			}
			sum += exp.Must(strconv.Atoi(string(row[rg[0]:rg[1]])))
		}

	}

	fmt.Println(sum)
}
