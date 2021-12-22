package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type point struct{ x, y, z int }

func main() {
	f, err := os.Open("inputs/22.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	cubes := make(map[point]struct{})
	for {
		var op string
		var minp, maxp point
		_, err := fmt.Fscanf(f, "%s x=%d..%d,y=%d..%d,z=%d..%d", &op, &minp.x, &maxp.x, &minp.y, &maxp.y, &minp.z, &maxp.z)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Couldn't scan: %v", err)
		}
		minp.x = max(-50, minp.x)
		maxp.x = min(50, maxp.x)
		minp.y = max(-50, minp.y)
		maxp.y = min(50, maxp.y)
		minp.z = max(-50, minp.z)
		maxp.z = min(50, maxp.z)

		var p point
		for p.x = minp.x; p.x <= maxp.x; p.x++ {
			for p.y = minp.y; p.y <= maxp.y; p.y++ {
				for p.z = minp.z; p.z <= maxp.z; p.z++ {
					if op == "on" {
						cubes[p] = struct{}{}
					} else {
						delete(cubes, p)
					}
				}
			}
		}
	}
	fmt.Println(len(cubes))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
