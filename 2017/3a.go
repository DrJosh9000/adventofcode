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
		log.Fatalf("Usage: 3a [number] (error was %v)", err)
	}
	
	step := []image.Point{
		{1, 0}, {0, -1}, {-1, 0}, {0, 1},
	}
	p := image.Pt(0, 0)
	spiral := map[image.Point]int{{}: 1}
	d := 0
	for i := 2; i <= input; i++ {
		p = p.Add(step[d])
		spiral[p] = i
		if _, adj := spiral[p.Add(step[(d+1)%4])]; !adj {
			d++
			d %= 4
		}
	}
	fmt.Println(p, abs(p.X) + abs(p.Y))
}
