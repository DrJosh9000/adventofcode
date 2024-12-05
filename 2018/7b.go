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

	active := make(map[rune]int)
	clock := 0
	for len(q) > 0 || len(active) > 0 {
		//fmt.Println(active)
		for u := range active {
			active[u]--
			if active[u] != 0 {
				continue
			}
			delete(active, u)
			for _, v := range g[u] {
				p[v]--
				if p[v] != 0 {
					continue
				}
				q = append(q, v)
			}
		}

		for len(q) > 0 && len(active) < 5 {
			sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })
			u := q[0]
			active[u] = int(u) - 4 // 60 + letter value
			q = q[1:]
		}

		clock++
	}

	fmt.Println(clock - 1)
}
