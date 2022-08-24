package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/DrJosh9000/exp"
)

type node struct {
	name string
	weight int
	above []string
}

func main() {
	nodes := make(map[string]*node)
	pred := make(map[string]int)
	exp.MustForEachLineIn("inputs/7.txt", func(line string) {
		var n node
		parts := strings.Split(line, " -> ")
		if _, err := fmt.Sscanf(parts[0], "%s (%d)", &n.name, &n.weight); err != nil {
			log.Fatalf("Couldn't scan first part: %v", err)
		}
		pred[n.name] = pred[n.name]
		nodes[n.name] = &n
		if len(parts) == 1 {
			return
		}
		n.above = strings.Split(parts[1], ", ")
		for _, k := range n.above {
			pred[k]++
		}
	})
	for k, n := range pred {
		if n == 0 {
			fmt.Println(k)
			return
		}
	}
}