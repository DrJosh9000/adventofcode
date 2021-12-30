package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/3.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	m := make(map[image.Point]int)
	for _, l := range strings.Split(string(f), "\n") {
		var p image.Point
		for _, token := range strings.Split(l, ",") {
			n, err := strconv.Atoi(token[1:])
			if err != nil {
				log.Fatalf("Couldn't atoi: %v", err)
			}
			switch token[0] {
			case 'U':
				for i := 0; i < n; i++ {
					p.Y++
					m[p]++
				}
			case 'D':
				for i := 0; i < n; i++ {
					p.Y--
					m[p]++
				}
			case 'L':
				for i := 0; i < n; i++ {
					p.X--
					m[p]++
				}
			case 'R':
				for i := 0; i < n; i++ {
					p.X++
					m[p]++
				}
			}
		}
	}

	md := math.MaxInt
	for p, x := range m {
		if x >= 2 {
			if d := abs(p.X) + abs(p.Y); d < md {
				md = d
			}
		}
	}
	fmt.Println(md)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
