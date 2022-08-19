package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Usage: 3b [number] (error was %v)", err)
	}
	
	step := []image.Point{
		{1, 0}, {0, -1}, {-1, 0}, {0, 1},
	}
	neigh := append(step, image.Pt(-1, -1), image.Pt(-1, 1), image.Pt(1, -1), image.Pt(1, 1))
	p := image.Pt(0, 0)
	spiral := map[image.Point]int{{}: 1}
	d := 0
	for {
		p = p.Add(step[d])
		for _, n := range neigh {
			spiral[p] += spiral[p.Add(n)]
		}
		if spiral[p] > input {
			fmt.Println(spiral[p])
			return
		}
		if _, adj := spiral[p.Add(step[(d+1)%4])]; !adj {
			d++
			d %= 4
		}
	}
}
