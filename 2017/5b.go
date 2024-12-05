package main

import (
	"fmt"
	"log"
	"strconv"

	"drjosh.dev/exp"
)

func main() {
	var offsets []int
	exp.MustForEachLineIn("inputs/5.txt", func(line string) {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Couldn't parse number: %v", err)
		}
		offsets = append(offsets, n)
	})

	steps, ip := 0, 0
	for ip < len(offsets) {
		steps++
		offset := offsets[ip]
		if offset >= 3 {
			offsets[ip]--
		} else {
			offsets[ip]++
		}
		ip += offset
	}
	fmt.Println(steps)
}
