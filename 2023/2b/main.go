package main

import (
	"fmt"
	"log"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2023
// Day 2, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("2023/inputs/2.txt") {
		_, ghs, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("Cut: %q", line)
		}

		var red, green, blue int

		for _, hand := range strings.Split(ghs, "; ") {
			for _, sp := range strings.Split(hand, ", ") {
				var n int
				var c string
				exp.Must(fmt.Sscanf(sp, "%d %s", &n, &c))

				switch c {
				case "red":
					red = max(red, n)
				case "green":
					green = max(green, n)
				case "blue":
					blue = max(blue, n)
				default:
					log.Fatalf("unknown colour %q", c)
				}
			}
		}

		sum += red * green * blue
	}
	fmt.Println(sum)
}
