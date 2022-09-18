package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2016
// Day 8

func main() {
	display := grid.Make[byte](6, 50)
	display.Fill('.')
	for _, line := range exp.MustReadLines("inputs/8.txt") {
		switch {
		case strings.HasPrefix(line, "rect"):
			var w, h int
			exp.Must(fmt.Sscanf(line, "rect %dx%d", &w, &h))
			display.FillRect(image.Rect(0, 0, w, h), '#')
		case strings.HasPrefix(line, "rotate row"):
			var y, r int
			exp.Must(fmt.Sscanf(line, "rotate row y=%d by %d", &y, &r))
			c := make([]byte, 50)
			for i := range c {
				c[(i+r)%50] = display[y][i]
			}
			display[y] = c
		case strings.HasPrefix(line, "rotate column"):
			var x, r int
			exp.Must(fmt.Sscanf(line, "rotate column x=%d by %d", &x, &r))
			c := make([]byte, 6)
			for i := range c {
				c[(i+r)%6] = display[i][x]
			}
			for i := range c {
				display[i][x] = c[i]
			}
		}
	}
	h := grid.Freq(display)
	fmt.Println(h['#'])
	for _, row := range display {
		fmt.Println(string(row))
	}
}
