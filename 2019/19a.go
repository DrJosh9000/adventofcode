package main

import (
	"fmt"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	ovm := intcode.ReadProgram("inputs/19.txt")

	test := func(x, y int) bool {
		vm := ovm.Copy()
		in, out := make(chan int, 2), make(chan int, 1)
		in <- x
		in <- y
		close(in)
		vm.Run(in, out)
		return <-out == 1
	}

	count := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if test(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}
