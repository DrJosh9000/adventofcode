package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("inputs/22.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	on := make(map[box]struct{})
lineLoop:
	for {
		var op string
		var r box
		_, err := fmt.Fscanf(f, "%s x=%d..%d,y=%d..%d,z=%d..%d", &op, &r.p.x, &r.q.x, &r.p.y, &r.q.y, &r.p.z, &r.q.z)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Couldn't scan: %v", err)
		}

		r.q.x++
		r.q.y++
		r.q.z++

		// Go through all the existing boxes, looking for overlap.
		for s := range on {
			o := r.intersect(s) // o for overlap
			if o.empty() {
				continue
			}
			// if r is fully contained within s, and r is on, then skip -
			// we'll just be removing and adding a box for no reason, and no
			// other boxes in the on set will overlap.
			if o == r && op == "on" {
				continue lineLoop
			}

			// the combination of r and s is nontrivial.
			// remove s from the set and replace it with boxes that make up
			// (s-o).
			delete(on, s)
			xc := []int{s.p.x, o.p.x, o.q.x, s.q.x}
			yc := []int{s.p.y, o.p.y, o.q.y, s.q.y}
			zc := []int{s.p.z, o.p.z, o.q.z, s.q.z}
			sort.Ints(xc)
			sort.Ints(yc)
			sort.Ints(zc)
			for i := range xc[1:] {
				for j := range yc[1:] {
					for k := range zc[1:] {
						t := box{
							p: point{xc[i], yc[j], zc[k]},
							q: point{xc[i+1], yc[j+1], zc[k+1]},
						}
						if t == o || t.empty() {
							continue
						}
						on[t] = struct{}{}
					}
				}
			}
		}
		// finally, if the new rectangle is on, add it
		if op == "on" {
			on[r] = struct{}{}
		}
	}

	vol := 0
	for r := range on {
		vol += r.volume()
	}
	fmt.Println(vol)
}

type point struct{ x, y, z int }

func (p point) in(r box) bool {
	return r.p.x <= p.x && p.x < r.q.x && r.p.y <= p.y && p.y < r.q.y && r.p.z <= p.z && p.z < r.q.z
}

type box struct{ p, q point }

func (r box) volume() int {
	return (r.q.x - r.p.x) * (r.q.y - r.p.y) * (r.q.z - r.p.z)
}

func (r box) empty() bool {
	return r.p.x >= r.q.x || r.p.y >= r.q.y || r.p.z >= r.q.z
}

func (r box) intersect(s box) box {
	if r.p.x < s.p.x {
		r.p.x = s.p.x
	}
	if r.p.y < s.p.y {
		r.p.y = s.p.y
	}
	if r.p.z < s.p.z {
		r.p.z = s.p.z
	}
	if r.q.x > s.q.x {
		r.q.x = s.q.x
	}
	if r.q.y > s.q.y {
		r.q.y = s.q.y
	}
	if r.q.z > s.q.z {
		r.q.z = s.q.z
	}
	if r.empty() {
		return box{}
	}
	return r
}
