package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 15, part a

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

	const row = 2000000
	poss := make(algo.Set[int])

	for _, s := range sensors {
		r := algo.L1(s.p.Sub(s.b))
		d := algo.Abs(row - s.p.Y)
		// Area intersects line at all?
		rd := r - d
		if rd < 0 {
			continue
		}

		for x := s.p.X - rd; x <= s.p.X+rd; x++ {
			poss.Insert(x)
		}
	}

	for _, s := range sensors {
		if s.b.Y == row {
			delete(poss, s.b.X)
		}
	}

	fmt.Println(len(poss))
}
