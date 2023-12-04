package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 4, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("2023/inputs/4.txt") {
		_, rest, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("no colon in %q", line)
		}
		wins, haves, ok := strings.Cut(rest, " | ")
		if !ok {
			log.Fatalf("no pipe in %q", line)
		}
		winm := make(algo.Set[int])
		for _, ws := range strings.Fields(wins) {
			winm.Insert(exp.Must(strconv.Atoi(ws)))
		}
		x := 0
		for _, hs := range strings.Fields(haves) {
			h := exp.Must(strconv.Atoi(hs))
			if !winm.Contains(h) {
				continue
			}
			if x == 0 {
				x = 1
			} else {
				x *= 2
			}
		}
		sum += x
	}
	fmt.Println(sum)
}
