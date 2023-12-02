package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 2, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("2023/inputs/2.txt") {
		gid, ghs, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("Cut: %q", line)
		}

		var id int
		exp.Must(fmt.Sscanf(gid, "Game %d", &id))

		possible := true

	handLoop:
		for _, hand := range strings.Split(ghs, "; ") {
			for _, sp := range strings.Split(hand, ", ") {
				var n int
				var c string
				exp.Must(fmt.Sscanf(sp, "%d %s", &n, &c))

				switch c {
				case "red":
					if n > 12 {
						possible = false
						break handLoop
					}

				case "green":
					if n > 13 {
						possible = false
						break handLoop
					}

				case "blue":
					if n > 14 {
						possible = false
						break handLoop
					}

				default:
					log.Fatalf("unknown colour %q", c)
				}
			}
		}

		if possible {
			sum += id
		}
	}
	fmt.Println(sum)
}
