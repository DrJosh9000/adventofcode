package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/parse"
)

// Advent of Code 2016
// Day 3, part b

func main() {
	var col [][]int
	for _, line := range exp.MustReadLines("inputs/3.txt") {
		tri := exp.Must(parse.Ints(line))
		for i := range tri {
			col[i] = append(col[i], tri[i])
		}
	}

	count := 0
	for _, c := range col {
	triLoop:
		for n := 0; n < len(c); n += 3 {
			tri := c[n : n+3]
			for i := range tri {
				for j := range tri {
					if i == j {
						continue
					}
					k := 3 - (i + j)
					if tri[k] >= tri[i]+tri[j] {
						continue triLoop
					}
				}
			}
			count++
		}
	}
	fmt.Println(count)
}
