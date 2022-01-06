package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/11.txt")
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)

	hull := make(map[image.Point]int)
	p, o := image.Pt(0, 0), image.Pt(0, 1)
commLoop:
	for {
		select {
		case c, ok := <-out:
			if !ok {
				// done
				break commLoop
			}
			hull[p] = c
			if <-out == 0 {
				o = image.Pt(-o.Y, o.X)
			} else {
				o = image.Pt(o.Y, -o.X)
			}
			p = p.Add(o)
		case in <- hull[p]:
			// nop
		}
	}
	fmt.Println(len(hull))
}
