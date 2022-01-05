package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("inputs/8.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	const h, w = 6, 25
	image := make([][]byte, h)
	for i := range image {
		image[i] = bytes.Repeat([]byte{'2'}, w)
	}

	for i, b := range f {
		if image[(i/w)%h][i%w] != '2' {
			continue
		}
		switch b {
		case '0':
			b = ' '
		case '1':
			b = '#'
		}
		image[(i/w)%h][i%w] = b
	}

	for _, l := range image {
		fmt.Println(string(l))
	}
}
