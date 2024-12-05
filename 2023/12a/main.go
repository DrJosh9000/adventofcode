package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 12, part a

const inputPath = "2023/inputs/12.txt"

func main() {
	lines := exp.MustReadLines(inputPath)

	fmt.Println(algo.Sum(algo.Map(lines, func(line string) int {
		springs, broks, ok := strings.Cut(line, " ")
		if !ok {
			log.Fatalf("couldn't cut line %q", line)
		}
		brok := algo.MustMap(strings.Split(broks, ","), strconv.Atoi)

		springs = "." + springs + "."
		pattern := []byte{'.'}
		for _, b := range brok {
			for n := 0; n < b; n++ {
				pattern = append(pattern, '#')
			}
			pattern = append(pattern, '.')
		}
		n1 := len(pattern) - 1

		count := make([]int, len(pattern))
		count[0] = 1

		for _, c := range []byte(springs) {
			nc := make([]int, len(pattern))

			for j, p := range pattern {
				switch c {
				case p, '?':
					if j == 0 {
						nc[j] = count[j]
						break
					}
					switch p {
					case '#':
						nc[j] = count[j-1]
					case '.':
						nc[j] = count[j] + count[j-1]
					}
				}
			}

			count = nc
		}
		//fmt.Printf("%s: %d\n", line, count)
		return count[n1]
	})))
}
