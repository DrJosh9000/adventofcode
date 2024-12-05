package main

import (
	"fmt"
	"log"

	"drjosh.dev/exp"
)

type nanobot struct {
	x, y, z, r int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(v, w nanobot) int {
	return abs(v.x - w.x) + abs(v.y - w.y) + abs(v.z - w.z)
}

func main() {
	var bots []nanobot
	maxr, best := -1, -1
	exp.MustForEachLineIn("inputs/23.txt", func(line string) {
		var bot nanobot
		if _, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &bot.x, &bot.y, &bot.z, &bot.r); err != nil {
			log.Fatalf("Couldn't scan line: %v", err)
		}
		if bot.r > maxr {
			maxr = bot.r
			best = len(bots)
		}
		bots = append(bots, bot)
	})

	count := 0
	for _, b := range bots {
		if dist(b, bots[best]) <= maxr {
			count++
		}
	}
	fmt.Println(count)
}
