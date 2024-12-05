package main

import (
	"fmt"
	"log"
	"sort"

	"drjosh.dev/exp"
)

func main() {
	// Topological sort
	g := make(map[rune][]rune)
	p := make(map[rune]int)
	exp.MustForEachLineIn("inputs/7.txt", func(line string) {
		var u, v rune
		if _, err := fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &u, &v); err != nil {
			log.Fatalf("Couldn't parse line %q: %v", line, err)
		}
		g[u] = append(g[u], v)
		p[u] = p[u] // NB: if u \notin p, p[u] == 0
		p[v]++
	})

	var q []rune
	for u, n := range p {
		if n == 0 {
			q = append(q, u)
		}
	}
	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })
		u := q[0]
		fmt.Printf("%c", u)
		q = q[1:]
		for _, v := range g[u] {
			p[v]--
			if p[v] == 0 {
				q = append(q, v)
			}
		}
	}
	fmt.Println()
}
