package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

func main() {
	nums := exp.MustReadInts("inputs/1.txt", "\n")
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
