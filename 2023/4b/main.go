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
// Day 4, part b

func main() {
	lines := exp.MustReadLines("2023/inputs/4.txt")
	table := make([]int, len(lines)+1)
	for i := range table {
		table[i] = 1
	}

	sum := 0
	for _, line := range lines {
		cardn, rest, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("no colon in %q", line)
		}
		wins, haves, ok := strings.Cut(rest, " | ")
		if !ok {
			log.Fatalf("no pipe in %q", line)
		}
		cardnum := exp.Must(strconv.Atoi(strings.Fields(cardn)[1]))
		winm := make(algo.Set[int])
		for _, ws := range strings.Fields(wins) {
			winm.Insert(exp.Must(strconv.Atoi(ws)))
		}
		matches := 0
		for _, hs := range strings.Fields(haves) {
			h := exp.Must(strconv.Atoi(hs))
			if !winm.Contains(h) {
				continue
			}
			matches++
		}
		sum += table[cardnum]
		for matches > 0 {
			table[cardnum+matches] += table[cardnum]
			matches--
		}
	}
	fmt.Println(sum)
}
