package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2016
// Day 3, part a

func main() {
	count := 0
lineLoop:
	for _, line := range exp.MustReadLines("inputs/3.txt") {
		tri := exp.Must(algo.MapOrErr(strings.Fields(line), strconv.Atoi))
		for i := range tri {
			for j := range tri {
				if i == j {
					continue
				}
				k := 3 - (i + j)
				if tri[k] >= tri[i]+tri[j] {
					continue lineLoop
				}
			}
		}
		count++
	}
	fmt.Println(count)
}
