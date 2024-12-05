package main

import (
	"fmt"

	"drjosh.dev/exp"
)

func main() {
	sum := 0
	exp.MustForEachLineIn("inputs/1.txt", func(line string) {
		for i := range line {
			if line[i] == line[(i+1)%len(line)] {
				sum += int(line[i] - '0')
			}
		}
	})
	fmt.Println(sum)
}