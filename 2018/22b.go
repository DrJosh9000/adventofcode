package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"os"

	"drjosh.dev/exp/algo"
)

func main() {
	f, err := os.Open("inputs/22.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	var depth int
	var target image.Point
	if _, err := fmt.Fscanf(f, "depth: %d\ntarget: %d,%d\n", &depth, &target.X, &target.Y); err != nil {
		log.Fatalf("Couldn't scan input: %v", err)
	}

	const (
		mod  = 20183
		xmul = 16807
		ymul = 48271
	)
	left, up := image.Pt(-1, 0), image.Pt(0, -1)
	erocache := map[image.Point]int{
		{0, 0}: depth % mod,
		target: depth % mod,
	}
	var erosion func(p image.Point) int
	erosion = func(p image.Point) int {
		if e, ok := erocache[p]; ok {
			return e
		}
		switch {
		case p.X < 0 || p.Y < 0:
			return -1
		case p.X == 0:
			erocache[p] = (p.Y*ymul + depth) % mod
		case p.Y == 0:
			erocache[p] = (p.X*xmul + depth) % mod
		default:
			erocache[p] = (erosion(p.Add(left))*erosion(p.Add(up)) + depth) % mod
		}
		return erocache[p]
	}

	neighs := []image.Point{left, up, image.Pt(1, 0), image.Pt(0, 1)}
	const (
		neither = 0
		torch   = 1
		gear    = 2
	)
	// ttype  tool   othertool
	// 0      1      2
	// 0      2      1
	// 1      0      2
	// 1      2      0
	// 2      0      1
	// 2      1      0
	// so othertool = 3 - (ttype + tool)
	type state struct {
		pos  image.Point
		tool int
	}
	start := state{tool: torch}
	want := state{pos: target, tool: torch}
	algo.Dijkstra(start, func(st state, dist int) (map[state]int, error) {
		//fmt.Printf("%+v\n", st)
		if st == want {
			fmt.Println(dist)
			return nil, errors.New("all done")
		}
		ttype := erosion(st.pos) % 3
		// consider changing tool for cost 7
		next := map[state]int{
			{pos: st.pos, tool: 3 - (ttype + st.tool)}: 7,
		}
		// consider moving to a compatible neighbour for cost 1
		for _, d := range neighs {
			q := st.pos.Add(d)
			ttype := erosion(q) % 3
			if ttype < 0 || ttype == st.tool {
				continue
			}
			next[state{pos: q, tool: st.tool}] = 1
		}
		return next, nil
	})
}
