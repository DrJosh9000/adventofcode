package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DrJosh9000/exp"
)

func main() {
	var nums []int
	exp.MustForEachLineIn("inputs/1.txt", func(line string) {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Parsing line: %v", err)
		}
		nums = append(nums, n)
	})
	dupes := make(map[int]bool)
	total := 0
	for {
		for _, n := range nums {
			total += n
			if dupes[total] {
				fmt.Println(total)
				return
			}
			dupes[total] = true
		}
	}
}