package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("inputs/17.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}
	defer f.Close()
	var target image.Rectangle
	if _, err := fmt.Fscanf(f, "target area: x=%d..%d, y=%d..%d", &target.Min.X, &target.Max.X, &target.Min.Y, &target.Max.Y); err != nil {
		log.Fatalf("Couldn't parse file: %v", err)
	}

	minvx := math.MaxInt
	for vx := 1; ; vx++ {
		t := vx * (vx + 1) / 2
		if t < target.Min.X {
			continue
		}
		if t > target.Max.X {
			break
		}
		if t < minvx {
			minvx = vx
		}
	}

	count := 0
	for ivy := target.Min.Y; ivy <= -target.Min.Y; ivy++ {
		for ivx := minvx; ivx <= target.Max.X; ivx++ {
			x, y := 0, 0
			vx, vy := ivx, ivy
			for {
				x += vx
				y += vy
				if vx > 0 {
					vx--
				}
				vy--
				if x > target.Max.X || y < target.Min.Y {
					break // missed
				}
				if x >= target.Min.X && x <= target.Max.X && y >= target.Min.Y && y <= target.Max.Y {
					count++
					break
				}
			}
		}
	}

	log.Print(count)
}
