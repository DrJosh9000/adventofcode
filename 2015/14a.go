package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2015
// Day 14, part a

func main() {
	type deer struct {
		name    string
		speed   int
		on, off int
		dist    int
	}

	var deers []*deer
	for _, line := range exp.MustReadLines("inputs/14.txt") {
		var d deer
		exp.Must(fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &d.name, &d.speed, &d.on, &d.off))
		deers = append(deers, &d)
	}

	const T = 2503
	maxd := 0
	for _, d := range deers {
		p := d.on + d.off
		on := (T / p) * d.on
		on += algo.Min(T%p, d.on)
		d.dist = on * d.speed
		if d.dist > maxd {
			maxd = d.dist
		}
	}

	fmt.Println(maxd)
}
