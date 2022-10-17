package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2015
// Day 6, part b

func main() {
	g := grid.Make[int](1000, 1000)
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
			g.MapRect(r, func(x int) int { return x + 1 })
		case strings.HasPrefix(line, "turn off"):
			g.MapRect(r, func(x int) int {
				if x <= 0 {
					return 0
				}
				return x - 1
			})
		case strings.HasPrefix(line, "toggle"):
			g.MapRect(r, func(x int) int { return x + 2 })
		}
	}

	sum := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			sum += g[y][x]
		}
	}
	fmt.Println(sum)
}
