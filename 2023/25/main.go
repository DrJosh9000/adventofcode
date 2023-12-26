package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 25

const inputPath = "2023/inputs/25.txt"

func main() {
	lines := exp.MustReadLines(inputPath)

	type edge struct{ u, v string }
	graph := make(map[string]algo.Set[string])
	var edges []edge

	for _, line := range lines {
		u, vs := exp.MustCut(line, ": ")
		for _, v := range strings.Split(vs, " ") {
			graph[u] = graph[u].Insert(v)
			graph[v] = graph[v].Insert(u)
			edges = append(edges, edge{u, v})
		}
	}
	// fmt.Println("|V| =", len(graph))
	// fmt.Println("|E| =", len(edges))

	for {
		perm := rand.Perm(len(edges))
		trees := make(algo.DisjointSets[string])
		var chosen []edge
		var last int
		for l, i := range perm {
			e := edges[i]
			if trees.Find(e.u) != trees.Find(e.v) {
				last = l
				chosen = append(chosen, e)
				trees.Union(e.u, e.v)
			}
		}

		cut := make(algo.DisjointSets[string])
		for _, c := range chosen[:len(chosen)-1] {
			cut.Union(c.u, c.v)
		}

		sets := cut.Sets()
		if len(sets) != 2 {
			// log.Printf("wrong number of sets %d != 2", len(sets))
			continue
		}

		count := 0
		for _, i := range perm[last:] {
			e := edges[i]
			if cut.Find(e.u) != cut.Find(e.v) {
				count++
			}
		}
		if count != 3 {
			// log.Printf("found cutset of %d edges", count)
			continue
		}

		prod := 1
		for _, set := range sets {
			prod *= len(set)
		}
		fmt.Println(prod)
		return
	}
}
