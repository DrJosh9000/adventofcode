package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/3.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var words []string
	for {
		var s string
		_, err := fmt.Fscan(f, &s)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't scan: %v", err)
		}
		words = append(words, s)
	}

	var gc [12]int
	for _, s := range words {
		for i, c := range s {
			if c == '1' {
				gc[i]++
			}
		}
	}
	gamma, epsilon := 0, 0
	for _, n := range gc {
		gamma *= 2
		epsilon *= 2
		if 2*n > len(words) {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	fmt.Println(gamma * epsilon)
}
