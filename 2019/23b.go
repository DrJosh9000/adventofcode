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
	var nmu sync.Mutex
	var nat image.Point
	lny := 0
	var actmu sync.Mutex
	active := 0x3ffffffffffff

	post := func(o int, p image.Point) {
		//fmt.Printf("%v -> %d\n", p, o)
		if o == 255 {
			nmu.Lock()
			nat = p
			nmu.Unlock()
			return
		}
		inmu[o].Lock()
		inqs[o] = append(inqs[o], p)
		inmu[o].Unlock()
	}

	activity := func(i int, a bool) {
		actmu.Lock()
		if a {
			active |= 1 << i
		} else {
			active &^= 1 << i
			if active == 0 {
				nmu.Lock()
				if nat != image.ZP {
					if nat.Y == lny {
						fmt.Println(lny)
						os.Exit(0)
					}
					post(0, nat)
					active |= 1
					lny = nat.Y
					nat = image.ZP
				}
				nmu.Unlock()
			}
		}
		actmu.Unlock()
	}

	peek := func(i int) image.Point {
		p := image.Point{-1, -1}
		inmu[i].Lock()
		act := len(inqs[i]) > 0
		if act {
			p = inqs[i][0]
		}
		inmu[i].Unlock()
		return p
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
					post(o, image.Pt(<-out, <-out))
					activity(i, true)
				case in <- p.X:
					activity(i, p.Y != -1)
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
