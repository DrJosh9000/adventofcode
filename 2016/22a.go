package main

import (
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 22, part a

func main() {
	type node struct {
		size, used, avail, usep int
	}
	nodes := make(map[image.Point]node)
	for _, line := range exp.MustReadLines("inputs/22.txt") {
		if !strings.HasPrefix(line, "/dev") {
			continue
		}
		var n node
		var p image.Point
		exp.Must(fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%", &p.X, &p.Y, &n.size, &n.used, &n.avail, &n.usep))
		nodes[p] = n
	}

	viable := 0
	for p, n := range nodes {
		if n.used == 0 {
			continue
		}
		for q, m := range nodes {
			if p == q {
				continue
			}
			if n.used <= m.avail {
				//fmt.Printf("can move %dT\tfrom %v\tto %v\n", n.used, p, q)
				viable++
			}
		}
	}
	fmt.Println(viable)
}
