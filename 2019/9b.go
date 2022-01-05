package main

import (
	"fmt"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/9.txt")
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)
	in <- 2
	for x := range out {
		fmt.Println(x)
	}
}
