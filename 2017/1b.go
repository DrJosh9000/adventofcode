package main

import (
	"fmt"

	"drjosh.dev/exp"
)

func main() {
	sum := 0
	exp.MustForEachLineIn("inputs/1.txt", func(line string) {
		offset := len(line) / 2
		for i := range line {
			if line[i] == line[(i+offset)%len(line)] {
				sum += int(line[i] - '0')
			}
		}
	})
	fmt.Println(sum)
}
