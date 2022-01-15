package main

import (
	"fmt"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/21.txt")
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)

	input := `NOT T T
AND A T
AND B T
AND C T
NOT T J
AND D J
WALK

`
	inpos := 0

	for {
		select {
		case x, ok := <-out:
			if !ok {
				return
			}
			if x < 128 {
				fmt.Printf("%c", x)
			} else {
				fmt.Println(x)
			}
		case in <- int(input[inpos]):
			inpos++
		}
	}
}
