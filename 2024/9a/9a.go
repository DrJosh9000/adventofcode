package main

import (
	_ "embed"
	"fmt"
)

//go:embed inputs/9.txt
var input string

func main() {
	var blocks []int
	id := 0
	for i, c := range input {
		if c < '0' || c > '9' {
			break
		}
		switch i % 2 {
		case 0:
			for range c - '0' {
				blocks = append(blocks, id)
			}
			id++

		case 1:
			for range c - '0' {
				blocks = append(blocks, -1)
			}
		}
	}

	for i := 0; i < len(blocks); i++ {
		if blocks[i] >= 0 {
			continue
		}
		for blocks[len(blocks)-1] < 0 {
			blocks = blocks[:len(blocks)-1]
		}
		if i >= len(blocks) {
			break
		}
		blocks[i] = blocks[len(blocks)-1]
		blocks = blocks[:len(blocks)-1]
	}

	cksum := 0
	for i, id := range blocks {
		cksum += i * id
	}

	fmt.Println(cksum)
}
