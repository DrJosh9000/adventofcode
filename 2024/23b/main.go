package main

import (
	"fmt"
	"slices"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 23, part b

const inputPath = "2024/inputs/23.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sets := make(map[string]algo.Set[string])
	for _, line := range lines {
		u, v := exp.MustCut(line, "-")
		sets[u] = sets[u].Insert(v)
		sets[v] = sets[v].Insert(u)
	}

	// largest clique is NP hard, right?
	cliques := make(algo.Set[string])
	for u, ua := range sets {
		for v := range ua {
			cl := []string{u, v}
			slices.Sort(cl)
			cliques.Insert(strings.Join(cl, ","))
		}
	}

	for u, ua := range sets {
		for cl := range cliques {
			cv := strings.Split(cl, ",")
			if slices.Contains(cv, u) {
				continue
			}
			if ua.ContainsAll(cv...) {
				cv2 := append(slices.Clone(cv), u)
				slices.Sort(cv2)
				cliques.Insert(strings.Join(cv2, ","))
			}
		}
	}

	largest := ""
	for cl := range cliques {
		if len(cl) > len(largest) {
			largest = cl
		}
	}

	fmt.Println(largest)
}
