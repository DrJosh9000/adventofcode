package main

import (
	"image"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/13.txt")
	vm[0] = 2
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)

	screen := make(map[image.Point]int)
	for x := range out {
		p := image.Pt(x, <-out)
		screen[p] = <-out
	}

	// WIP
}
