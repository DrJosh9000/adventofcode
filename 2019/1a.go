package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/1.txt")
	if err != nil {
		log.Fatalf("Couldn't open input file: %v", err)
	}
	defer f.Close()

	sum := 0
	for {
		var x int
		_, err := fmt.Fscan(f, &x)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Couldn't scan int: %v", err)
		}
		sum += (x / 3) - 2
	}
	fmt.Println(sum)
}
