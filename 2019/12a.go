package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/12.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	var pos, vel []vec
	for i := 0; i < 4; i++ {
		var p vec
		if _, err := fmt.Fscanf(f, "<x=%d, y=%d, z=%d>\n", &p[0], &p[1], &p[2]); err != nil {
			log.Fatalf("Couldn't scan position: %v", err)
		}
		pos = append(pos, p)
		vel = append(vel, vec{})
	}

	for i := 0; i < 1000; i++ {
		for j, x := range pos {
			for k, y := range pos[j:] {
				d := x.add(y.neg()).sign() // sign(x - y)
				vel[j] = vel[j].add(d.neg())
				vel[j+k] = vel[j+k].add(d)
			}
		}
		for j := range pos {
			pos[j] = pos[j].add(vel[j])
		}
	}
	total := 0
	for i := range pos {
		total += pos[i].norm() * vel[i].norm()
	}
	fmt.Println(total)
}

type vec [3]int

func (v vec) add(w vec) vec {
	return [3]int{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

func (v vec) neg() vec {
	return [3]int{-v[0], -v[1], -v[2]}
}

func (v vec) sign() vec {
	return [3]int{sign(v[0]), sign(v[1]), sign(v[2])}
}

func (v vec) norm() int {
	return abs(v[0]) + abs(v[1]) + abs(v[2])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}
