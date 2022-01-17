package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/25.txt")
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)

	r := bufio.NewReader(os.Stdin)
	for {
		var last9 []byte
		prompt := false
		for o := range out {
			fmt.Printf("%c", o)
			last9 = append(last9, byte(o))
			if len(last9) > 9 {
				last9 = last9[len(last9)-9:]
			}
			if string(last9) == "Command?\n" {
				prompt = true
				break
			}
		}
		if !prompt {
			return
		}
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("Couldn't read: %v", err)
		}
		for _, b := range i {
			in <- int(b)
		}
	}
}
