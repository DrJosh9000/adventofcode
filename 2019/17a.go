package main

import (
	"fmt"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	vm := intcode.ReadProgram("inputs/17.txt")
	out := make(chan int)
	go vm.Run(nil, out)
	var camera [][]byte
	var line []byte
	for x := range out {
		if x == '\n' {
			if len(line) > 0 {
				camera = append(camera, line)
			}
			line = nil
			continue
		}
		line = append(line, byte(x))
	}
	if len(line) > 0 {
		camera = append(camera, line)
	}

	if false {
		for _, l := range camera {
			fmt.Println(string(l))
		}
	}

	h, w := len(camera), len(camera[0])
	sum := 0
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if camera[y][x] == '.' || camera[y-1][x] == '.' || camera[y+1][x] == '.' || camera[y][x-1] == '.' || camera[y][x+1] == '.' {
				continue
			}
			sum += x * y
		}
	}
	fmt.Println(sum)
}
