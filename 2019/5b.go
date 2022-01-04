package main

import (
	"fmt"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.VM{
		M:   intcode.ReadProgram("inputs/5.txt"),
		In:  make(chan int, 1),
		Out: make(chan int),
	}
	vm.In <- 5
	done := make(chan struct{})
	go func() {
		for x := range vm.Out {
			fmt.Println(x)
		}
		close(done)
	}()
	vm.Run()
	<-done
}
