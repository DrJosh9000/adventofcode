package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 24, part a

const inputPath = "2024/inputs/24.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	values := make(map[string]int)
	rdeps := make(map[string][]string)
	gates := make(map[string]gate)
	for _, line := range lines {
		var wire string
		var value int
		var gate gate
		switch {
		case exp.Smatchf(line, "%s %d", &wire, &value):
			values[strings.TrimSuffix(wire, ":")] = value
		case exp.Smatchf(line, "%s %s %s -> %s", &gate.x, &gate.op, &gate.y, &gate.z):
			gates[gate.z] = gate
			rdeps[gate.x] = append(rdeps[gate.x], gate.z)
			rdeps[gate.y] = append(rdeps[gate.y], gate.z)
		}
	}

	q := slices.Collect(maps.Keys(values))
	for len(q) > 0 {
		wire := q[0]
		q = q[1:]
		for _, rd := range rdeps[wire] {
			g := gates[rd]
			x, xok := values[g.x]
			y, yok := values[g.y]
			if !xok || !yok {
				continue
			}
			switch g.op {
			case "AND":
				values[g.z] = x & y
			case "OR":
				values[g.z] = x | y
			case "XOR":
				values[g.z] = x ^ y
			}
			q = append(q, g.z)
		}
	}

	n, b := 0, 1
	for zi := 0; ; zi++ {
		z := fmt.Sprintf("z%02d", zi)
		x, ok := values[z]
		if !ok {
			break
		}
		n += x * b
		b <<= 1
	}
	fmt.Println(n)
}

type gate struct {
	x, y, z, op string
}
