package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 15, part b

type sensor struct {
	p image.Point
	b image.Point
}

func main() {
	var sensors []sensor
	for _, line := range exp.MustReadLines("inputs/15.txt") {
		var s sensor
		exp.Must(fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.p.X, &s.p.Y, &s.b.X, &s.b.Y))
		sensors = append(sensors, s)
	}

	bounds := image.Rect(0, 0, 4000001, 4000001)

	for _, s := range sensors {
		r := algo.L1(s.p.Sub(s.b)) + 1

		// Walk around the circumference
		var d image.Point
		for d.X = -r; d.X <= r; d.X++ {
			adx := algo.Abs(d.X)
		coordLoop:
			for _, d.Y = range []int{r - adx, adx - r} {
				p := s.p.Add(d)
				if !p.In(bounds) {
					continue
				}
				for _, t := range sensors {
					if algo.L1(p.Sub(t.p)) < algo.L1(t.p.Sub(t.b)) {
						continue coordLoop
					}
				}
				fmt.Println(p.X*4000000 + p.Y)
				return
			}
		}
	}
}
