package main

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"

	"drjosh.dev/exp/algo"
)

func hash(input []byte) []byte {
	input = append(input, 17, 31, 73, 47, 23)

	circle := make([]byte, 256)
	for i := range circle {
		circle[i] = byte(i)
	}
	var pos, skip byte

	for round := 0; round < 64; round++ {
		for _, l := range input {
			for i := byte(0); i < l/2; i++ {
				j, k := pos+i, pos+l-i-1
				circle[j], circle[k] = circle[k], circle[j]
			}
			pos += l + skip
			skip++
		}
	}

	dense := make([]byte, 16)
	for i := 0; i < 16; i++ {
		dense[i] = algo.Foldl(circle[i*16:(i+1)*16], func(x, y byte) byte {
			return x ^ y
		})
	}
	return dense
}

func main() {
	input := os.Args[1]

	count := 0
	for i := 0; i < 128; i++ {
		for _, b := range hash([]byte(input + "-" + strconv.Itoa(i))) {
			count += bits.OnesCount8(b)
		}
	}
	fmt.Println(count)
}
