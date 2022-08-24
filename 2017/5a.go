package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DrJosh9000/exp"
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
		ip, offsets[ip] = ip+offsets[ip], offsets[ip]+1
	}
	fmt.Println(steps)
}