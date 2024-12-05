package main

import (
	"fmt"
	"image"
	"log"

	"drjosh.dev/exp"
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

outerLoop:
	for i, c := range claims {
		for j, d := range claims {
			if i == j {
				continue
			}
			if c.b.Overlaps(d.b) {
				continue outerLoop
			}
		}
		fmt.Println(c.id)
		return
	}
}
