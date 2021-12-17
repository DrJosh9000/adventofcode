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

	// So. If x += vx, and vx--, that means the target x position will be the
	// (vx)th triangular number, and there will be at least vx steps.
	// Find the range of triangular numbers in the target.
	minvx, maxvx := math.MaxInt, math.MinInt
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
		if t > maxvx {
			maxvx = vx
		}
	}
	//log.Print(minvx, maxvx)

	// Now for vy. Whatever positive vy we use, the probe will be back at y=0
	// after ~2*vy steps, with the same vertical speed as at launch, but
	// downward.
	// But that means we can have a huge initial vy that crashes back down
	// from y=0 to min Y in a single step! Any faster and we will always miss,
	// so that's the upper bound for the search.
	maxy := 0
	for ivy := 0; ivy <= -target.Min.Y; ivy++ {
		for ivx := minvx; ivx <= maxvx; ivx++ {
			my := 0
			x, y := 0, 0
			vx, vy := ivx, ivy
			for {
				x += vx
				y += vy
				if vx > 0 {
					vx--
				}
				vy--
				if y > my {
					my = y
				}
				if x > target.Max.X || y < target.Min.Y {
					break // missed
				}
				if x >= target.Min.X && x <= target.Max.X && y >= target.Min.Y && y <= target.Max.Y {
					if my > maxy {
						maxy = my
					}
					break
				}
			}
		}
	}

	log.Print(maxy)
}
