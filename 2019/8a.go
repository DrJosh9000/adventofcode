package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.ReadFile("inputs/8.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	const h, w = 6, 25
	ls := h * w
	var stats [][]int
	for i, d := range f {
		l := i / ls
		for l >= len(stats) {
			stats = append(stats, []int{0, 0, 0})
		}
		stats[l][d-'0']++
	}

	best := math.MaxInt
	var bestl []int
	for _, l := range stats {
		if l[0] < best {
			best, bestl = l[0], l
		}
	}
	fmt.Println(bestl[1] * bestl[2])
}
