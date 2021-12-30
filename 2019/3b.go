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

	dm := []map[image.Point]int{
		make(map[image.Point]int),
		make(map[image.Point]int),
	}
	for j, l := range strings.Split(string(f), "\n") {
		var p image.Point
		d := 0
		relax := func() {
			d++
			if _, exists := dm[j][p]; !exists {
				dm[j][p] = d
			}

		}
		for _, token := range strings.Split(l, ",") {
			n, err := strconv.Atoi(token[1:])
			if err != nil {
				log.Fatalf("Couldn't atoi: %v", err)
			}
			switch token[0] {
			case 'U':
				for i := 0; i < n; i++ {
					p.Y++
					relax()
				}
			case 'D':
				for i := 0; i < n; i++ {
					p.Y--
					relax()
				}
			case 'L':
				for i := 0; i < n; i++ {
					p.X--
					relax()
				}
			case 'R':
				for i := 0; i < n; i++ {
					p.X++
					relax()
				}
			}
		}
	}

	md := math.MaxInt
	for p, d0 := range dm[0] {
		d1, found := dm[1][p]
		if !found {
			continue
		}
		if d := d0 + d1; d < md {
			md = d
		}

	}
	fmt.Println(md)
}
