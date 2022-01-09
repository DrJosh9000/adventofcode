package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/15.txt")

	d := map[int]image.Point{1: {0, -1}, 2: {0, 1}, 3: {-1, 0}, 4: {1, 0}}

	type state *intcode.VM
	area := map[image.Point]*intcode.VM{{}: vm}
	var oxygen image.Point
	q := []image.Point{{}}
	var p image.Point
	for len(q) > 0 {
		p, q = q[0], q[1:]
		st := area[p]
		for i, d := range d {
			r := p.Add(d)
			if _, known := area[r]; known {
				continue
			}
			vm := st.Copy()
			in, out := make(chan int, 1), make(chan int, 1)
			in <- i
			close(in)
			vm.Run(in, out)
			switch <-out {
			case 0:
				continue
			case 1:
				area[r] = vm
				q = append(q, r)
			case 2:
				oxygen = r
				area[r] = vm
				q = append(q, r)
			}
		}
	}

	best := 0
	dist := map[image.Point]int{oxygen: 0}
	q = []image.Point{oxygen}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		t := dist[p] + 1
		delete(area, p)
		for _, d := range d {
			r := p.Add(d)
			if _, conn := area[r]; !conn {
				continue
			}
			dist[r] = t
			if t > best {
				best = t
			}
			q = append(q, r)
		}
	}
	fmt.Println(best)
}
