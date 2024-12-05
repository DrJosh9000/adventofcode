package main

import (
	"fmt"
	"log"
	"strings"

	"drjosh.dev/exp"
)

type node struct {
	name        string
	weight, stw int
	aboveKey    []string
	above       []*node
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
		n.aboveKey = strings.Split(parts[1], ", ")
		for _, k := range n.aboveKey {
			pred[k]++
		}
	})
	for _, n := range nodes {
		for _, a := range n.aboveKey {
			n.above = append(n.above, nodes[a])
		}
	}

	var root *node
	for k, n := range pred {
		if n == 0 {
			root = nodes[k]
			break
		}
	}

	var sumWeights func(*node) int
	sumWeights = func(n *node) int {
		if n.stw != 0 {
			return n.stw
		}
		n.stw = n.weight
		for _, a := range n.above {
			n.stw += sumWeights(a)
		}
		return n.stw
	}

	sumWeights(root)

	var search func(*node) bool
	search = func(n *node) bool {
		if len(n.above) <= 1 {
			// trivially balanced
			return false
		}
		m := make(map[int]int)
		for _, a := range n.above {
			m[a.stw]++
		}
		if len(m) == 1 {
			// my subtowers are balanced
			return false
		}
		// hope len(m) != 2 ...
		var common int
		for w, x := range m {
			if x != 1 {
				common = w
			}
		}
		for _, a := range n.above {
			if a.stw == common {
				continue
			}
			if !search(a) {
				// a's subtowers are balanced but mine are not...
				// therefore it is a's fault
				fmt.Println(a.weight + common - a.stw)
				return true
			}
		}
		return true
	}

	search(root)
}
