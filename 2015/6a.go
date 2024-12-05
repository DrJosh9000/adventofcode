package main

import (
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2015
// Day 6, part a

// This is day 6. Plan accordingly.

func main() {
	g := grid.Make[bool](1000, 1000)
	for _, line := range exp.MustReadLines("inputs/6.txt") {
		f := strings.Fields(line)
		n1 := len(f) - 1
		var r image.Rectangle
		exp.Must(fmt.Sscanf(f[n1-2], "%d,%d", &r.Min.X, &r.Min.Y))
		exp.Must(fmt.Sscanf(f[n1], "%d,%d", &r.Max.X, &r.Max.Y))
		r.Max.X++
		r.Max.Y++
		switch {
		case strings.HasPrefix(line, "turn on"):
			g.FillRect(r, true)
		case strings.HasPrefix(line, "turn off"):
			g.FillRect(r, false)
		case strings.HasPrefix(line, "toggle"):
			g.MapRect(r, func(x bool) bool { return !x })
		}
	}

	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if g[y][x] {
				count++
			}
		}
	}
	fmt.Println(count)
}
