package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2015
// Day 14, part b

func main() {
	type deer struct {
		name    string
		speed   int
		on, off int
		dist    int
		points  int
	}

	var deers []*deer
	for _, line := range exp.MustReadLines("inputs/14.txt") {
		var d deer
		exp.Must(fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &d.name, &d.speed, &d.on, &d.off))
		deers = append(deers, &d)
	}

	const T = 2503
	for i := 0; i < 2503; i++ {
		maxd := 0
		var winners []int
		for j, d := range deers {
			p := d.on + d.off
			if i%p < d.on {
				d.dist += d.speed
			}
			if d.dist == maxd {
				winners = append(winners, j)
			}
			if d.dist > maxd {
				maxd = d.dist
				winners = []int{j}
			}
		}
		for _, j := range winners {
			deers[j].points++
		}
	}

	maxp := 0
	for _, d := range deers {
		if d.points > maxp {
			maxp = d.points
		}
	}
	fmt.Println(maxp)
}
