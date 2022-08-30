package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

func main() {
	fw := make(map[int]int)
	exp.MustForEachLineIn("inputs/13.txt", func(line string) {
		var d, r int
		exp.Must(fmt.Sscanf(line, "%d: %d", &d, &r))
		fw[d] = r
	})
	
	sev := 0
	for d, r := range fw {
		if d%((r-1)*2) == 0 {
			sev += d * r
		}
	}
	fmt.Println(sev)
}