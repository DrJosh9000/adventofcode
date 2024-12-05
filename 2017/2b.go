package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

func main() {
	sum := 0
	exp.MustForEachLineIn("inputs/2.txt", func(line string) {
		var row []int
		for _, ns := range strings.Fields(line) {
			n, err := strconv.Atoi(ns)
			if err != nil {
				log.Fatalf("Couldn't parse %q: %v", ns, err)
			}
			row = append(row, n)
		}

		for i := range row {
			for j := range row {
				if i == j {
					continue
				}
				if row[i]%row[j] == 0 {
					sum += row[i] / row[j]
					return
				}
			}
		}
	})
	fmt.Println(sum)
}
