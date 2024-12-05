package main

import (
	"fmt"

	"drjosh.dev/exp"
)

func main() {
	fw := make(map[int]int)
	exp.MustForEachLineIn("inputs/13.txt", func(line string) {
		var d, r int
		exp.Must(fmt.Sscanf(line, "%d: %d", &d, &r))
		fw[d] = r
	})
	
delayLoop:
	for delay := 1; ; delay++ {
		for d, r := range fw {
			if (d + delay) % ((r-1)*2) == 0 {
				continue delayLoop
			}
		}

		fmt.Println(delay)
		return
	}
}