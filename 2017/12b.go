package main

import (
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

func main() {
	d := make(algo.DisjointSets[string])
	exp.MustForEachLineIn("inputs/12.txt", func(line string) {
		arrow := strings.Split(line, " <-> ")
		src := arrow[0]
		for _, dst := range strings.Split(arrow[1], ", ") {
			d.Union(src, dst)
		}
	})
	
	groups := make(map[string]struct{})
	for x := range d {
		groups[d.Find(x)] = struct{}{}
	}
	fmt.Println(len(groups))
}