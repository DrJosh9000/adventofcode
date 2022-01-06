package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/13.txt")
	out := make(chan int)
	go vm.Run(nil, out)
	screen := make(map[image.Point]int)
	for x := range out {
		p := image.Pt(x, <-out)
		screen[p] = <-out
	}

	count := 0
	for _, id := range screen {
		if id == 2 {
			count++
		}
	}
	fmt.Println(count)
}
