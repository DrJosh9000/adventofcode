package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/15.txt")

	d := map[int]image.Point{1: {0, -1}, 2: {0, 1}, 3: {-1, 0}, 4: {1, 0}}

	type state struct {
		vm    *intcode.VM
		steps int
	}
	area := map[image.Point]state{{}: {vm: vm}}
	q := []image.Point{{}}
	var p image.Point
	for len(q) > 0 {
		p, q = q[0], q[1:]
		st := area[p]
		for i, d := range d {
			if _, known := area[p.Add(d)]; known {
				continue
			}
			vm := st.vm.Copy()
			in, out := make(chan int, 1), make(chan int, 1)
			in <- i
			close(in)
			vm.Run(in, out)
			switch <-out {
			case 0:
				continue
			case 1:
				area[p.Add(d)] = state{vm: vm, steps: st.steps + 1}
				q = append(q, p.Add(d))
			case 2:
				fmt.Println(st.steps + 1)
				return
			}
		}
	}
}
