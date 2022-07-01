 package main

import (
	"fmt"
	"image"
	"log"

	"github.com/DrJosh9000/exp"
)

type claim struct {
	id int
	b image.Rectangle
}

func main() {
	var claims []claim
	exp.MustForEachLineIn("inputs/3.txt", func(line string) {
		var id, x, y, w, h int
		if _, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h); err != nil {
			log.Fatalf("Scanning line: %v", err)
		}
		claims = append(claims, claim{
			id: id,
			b: image.Rect(x, y, x+w, y+h),
		})
	})
	doubles := make(map[image.Point]struct{})
	for i, c := range claims[:len(claims)-1] {
		for _, d := range claims[i+1:] {
			r := c.b.Intersect(d.b)
			if r.Empty() {
				continue
			}
			for x := r.Min.X; x < r.Max.X; x++ {
				for y := r.Min.Y; y < r.Max.Y; y++ {
					doubles[image.Pt(x, y)] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(doubles))
}