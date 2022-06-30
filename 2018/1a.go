package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DrJosh9000/exp"
)

func main() {
	var total int
	exp.MustForEachLineIn("inputs/1.txt", func(line string) {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Parsing line: %v", err)
		}
		total += n
	})
	fmt.Println(total)
}