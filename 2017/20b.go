package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

type particle struct {
	pos, vel, acc [3]int
}

func add(v, w [3]int) [3]int {
	return [3]int{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

func main() {
	swarm := make(algo.Set[*particle])
	chart := make(map[[3]int]algo.Set[*particle])

	exp.MustForEachLineIn("inputs/20.txt", func(line string) {
		var p particle
		exp.Must(fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p.pos[0], &p.pos[1], &p.pos[2], &p.vel[0], &p.vel[1], &p.vel[2], &p.acc[0], &p.acc[1], &p.acc[2]))
		swarm.Insert(&p)
		if chart[p.pos] == nil {
			chart[p.pos] = make(algo.Set[*particle])
		} 
		chart[p.pos].Insert(&p)
	})

	// Ehhh... just simulate 1000 ticks? Good enough for AoC!
	for t := 1; t < 1_000; t++ {
		for p := range swarm {
			delete(chart[p.pos], p)
			if len(chart[p.pos]) == 0 {
				delete(chart, p.pos)
			}
			p.vel = add(p.vel, p.acc)
			p.pos = add(p.pos, p.vel)
			if chart[p.pos] == nil {
				chart[p.pos] = make(algo.Set[*particle])
			} 
			chart[p.pos].Insert(p)
		}

		for v, s := range chart {
			if len(s) < 2 {
				continue
			}
			for p := range s {
				delete(swarm, p)
			}
			delete(chart, v)
		}
	}
	
	// A more involved idea... buggy:

	/*
		func (p *particle) posAt(t int) [3]int {
			// v(t) = v(t-1) + a
			// v(t) = v(0) + a*t  (clear)
			// p(t) = p(t-1) + v(t)
			//      = p(t-1) + v(0) + a*t
			// p(0) = p(0) + 0*v(0) + 0*1*a/2 = p(0)   (base case, trivial)
			// p(1) = p(0) + v(0) + a
			// p(2) = p(0) + v(0) + a + v(0) + 2a = p(0) + 2v(0) + 3a
			// p(3) = p(0) + 2v(0) + 3a + v(0) + 3a = p(0) + 3v(0) + 6a
			// RTP inductive case:
			// p(t-1) = p(0) + (t-1)*v(0) + t*(t-1)*a/2 => p(t) = p(0) + t*v(0) + t*(t+1)*a/2
			// Assume p(t-1) = p(0) + (t-1)*v(0) + t*(t-1)*a/2, then
			// p(t) = p(t-1) + v(t)
			//      = p(0) + (t-1)*v(0) + t*(t-1)*a/2 + v(0) + a*t
			//      = p(0) + t*v(0) + (t * a/2) * (t-1 + 2)
			//      = p(0) + t*v(0) + t*(t+1)*a/2.  By induction, etc. QED.
			tri := t * (t+1) / 2
			return [3]int{
				p.pos[0] + t * p.vel[0] + tri * p.acc[0],
				p.pos[1] + t * p.vel[1] + tri * p.acc[1],
				p.pos[2] + t * p.vel[2] + tri * p.acc[2],
			}
		}
	
		type spacetime struct {
			pos [3]int
			time int
		}

		// coordinate -> all particles that would collide at that coordinate
		collisions := make(map[spacetime]algo.Set[*particle])

		// For all pairs of particles, record the coordinates when they collide...
		for p := range swarm {
		pairLoop:
			for q := range swarm {
				if p == q {
					continue
				}

				// Find candidate times for collision based on each dimension.
				solns := make(algo.Set[int])
				for dim := 0; dim < 3; dim++ {
					// p(t) = p'(t) ==> p(0) + t*v(0) + t*(t+1)*a/2 = p'(0) + t*v'(0) + t*(t+1)*a'/2
					// 			    ==> 0 = (p(0) - p'(0)) + t*(v(0) - v'(0)) + t*(t+1)*(a-a')/2
					// 				==> 0 = Δp + t*(Δv) + t*(t+1)*Δa/2
					//              ==> 0 = Δp + t*(Δv + Δa/2) + t*t*Δa/2
					// Doubling yields an equation with the same solutions, and also integer coefficients:
					//                  0 = 2Δp + t*(2Δv + Δa) + t*t*Δa
					a := p.acc[dim] - q.acc[dim]
					b := 2*(p.vel[dim] - q.vel[dim]) + a
					c := 2*(p.pos[dim] - q.pos[dim])

					if a == 0 {
						// Reduces to linear equation.
						// 0 = Δp + t*Δv   ==>   t = -Δp / Δv = -c/b
						// b = 2Δv (because Δa = 0), and c = 2Δp
						if b == 0 {
							// Reduces to constants.
							if c != 0 { // a fixed distance apart
								continue pairLoop
							}
							// Two equal particles...
							solns.Insert(1)
							continue
						}
						if -c % b != 0 {
							// Collision happens, but not at an integer tick.
							continue pairLoop
						}
						// Collision happens at this one solution.
						solns.Insert(-c / b)
						continue
					}
					disc := b*b - 4*a*c
					if disc < 0 {
						// No real solutions on this axis - no collision.
						continue pairLoop
					}
					rdisc := sort.Search(disc, func(x int) bool { return x*x > disc }) - 1
					if rdisc * rdisc != disc {
						// disc is not a square - collision happens, but not at an integer tick.
						continue pairLoop
					}
					// Check that the collisions happen at integer ticks
					if (-b - rdisc) % (2*a) == 0 {
						solns.Insert((-b - rdisc) / 2*a)
					}
					if (-b + rdisc) % (2*a) == 0 {
						solns.Insert((-b + rdisc) / 2*a)
					}
				}

				// Check the candidate solutions for t to see if the positions align.
				for t := range solns {
					if t < 0 {
						continue
					}
					c := spacetime{p.posAt(t), t}
					if c.pos != q.posAt(t) {
						continue
					}
					if collisions[c] == nil {
						collisions[c] = make(algo.Set[*particle])
					}
					collisions[c].Insert(p, q)
				}
			}
		}

		// March forward, removing particles that collide from both
		// the swarm and from all future collisions.
		coords := algo.Keys(collisions)
		sort.Slice(coords, func(i, j int) bool {
			return coords[i].time < coords[j].time
		})
		for _, c := range coords {
			coll := collisions[c]
			for p := range coll {
				if !swarm.Contains(p) {
					delete(coll, p)
				}
			}
			if len(coll) < 2 {
				// A collision that would have happened here doesn't actually happen,
				// because one or more particles had earlier collisions.
				continue
			}
			for p := range coll {
				delete(swarm, p)
			}
		}
	*/

	fmt.Println(len(swarm))
}
