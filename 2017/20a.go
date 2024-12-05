package main

import (
	"fmt"
	"math"

	"drjosh.dev/exp"
)

type particle struct {
	pos, vel, acc [3]int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func norm(v [3]int) int {
	return abs(v[0]) + abs(v[1]) + abs(v[2])
}

func main() {
	var swarm []particle
	exp.MustForEachLineIn("inputs/20.txt", func(line string) {
		var p particle
		exp.Must(fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p.pos[0], &p.pos[1], &p.pos[2], &p.vel[0], &p.vel[1], &p.vel[2], &p.acc[0], &p.acc[1], &p.acc[2]))
		swarm = append(swarm, p)
	})

	var best int
	ban, bvn, bpn := math.MaxInt, math.MaxInt, math.MaxInt
	for i, p := range swarm {
		an := norm(p.acc)
		if an < ban {
			best = i
			ban, bvn, bpn = an, norm(p.vel), norm(p.pos)
			continue
		}
		if an > ban {
			continue
		}
		vn := norm(p.vel)
		if vn < bvn {
			best = i
			bvn, bpn = vn, norm(p.pos)
			continue
		}
		if vn > bvn {
			continue
		}
		pn := norm(p.pos)
		if pn < bpn {
			best = i
			bpn = pn
		}
	}

	fmt.Println(best)
}
