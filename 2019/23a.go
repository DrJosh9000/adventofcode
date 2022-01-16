package main

import (
	"fmt"
	"image"
	"os"
	"sync"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	ovm := intcode.ReadProgram("inputs/23.txt")

	vms := make([]*intcode.VM, 50)
	for i := range vms {
		vms[i] = ovm.Copy()
	}

	inmu := make([]sync.Mutex, 50)
	inqs := make([][]image.Point, 50)

	peek := func(i int) image.Point {
		p := image.Point{-1, -1}
		inmu[i].Lock()
		if len(inqs[i]) > 0 {
			p = inqs[i][0]
		}
		inmu[i].Unlock()
		return p
	}
	post := func(o int, p image.Point) {
		inmu[o].Lock()
		inqs[o] = append(inqs[o], p)
		inmu[o].Unlock()
	}
	pop := func(i int) {
		inmu[i].Lock()
		inqs[i] = inqs[i][1:]
		inmu[i].Unlock()
	}

	for i, vm := range vms {
		i, vm := i, vm
		in, out := make(chan int), make(chan int)
		go vm.Run(in, out)
		go func() {
			in <- i
			for {
				p := peek(i)
				select {
				case o := <-out:
					p.X, p.Y = <-out, <-out
					if o == 255 {
						fmt.Println(p.Y)
						os.Exit(0)
					}
					post(o, p)
				case in <- p.X:
					if p.Y != -1 {
						in <- p.Y
						pop(i)
					}
				}
			}
		}()
	}

	select {}
}
