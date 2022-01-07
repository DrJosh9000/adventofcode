package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/13.txt")
	vm[0] = 2
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)

	var score, joy, padx int
commLoop:
	for {
		select {
		case x, ok := <-out:
			if !ok {
				break commLoop
			}
			p := image.Pt(x, <-out)
			id := <-out
			if p == image.Pt(-1, 0) {
				score = id
				continue
			}
			switch id {
			case 4: // ball
				joy = sign(x - padx)
			case 3: // paddle
				padx = x
			}
		case in <- joy:
			//nop
		}
	}
	fmt.Println(score)
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
