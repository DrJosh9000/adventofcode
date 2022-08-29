package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

func main() {
	// Read the file, remove the trailing newline, convert to ints, append standard suffix.
	lengths := append(
		algo.Map(
			bytes.TrimSpace(exp.Must(os.ReadFile("inputs/10.txt"))), 
			func(b byte) int {
				return int(b)
			},
		), 
		17, 31, 73, 47, 23,
	)
	
	circle := make([]byte, 256)
	for i := range circle {
		circle[i] = byte(i)
	}
	pos, skip := 0, 0
	
	for round := 0; round < 64; round++ {
		for _, l := range lengths {
			for i := 0; i < l/2; i++ {
				j, k := (pos+i)%len(circle), (pos+l-i-1)%len(circle)
				circle[j], circle[k] = circle[k], circle[j]
			}
			pos += l + skip
			skip++
		}
	}
	
	dense := make([]byte, 16)
	for i := 0; i < 16; i++ {
		dense[i] = algo.Reduce(circle[i*16:(i+1)*16], func(x, y byte) byte {
			return x ^ y
		})
	}
	
	fmt.Printf("%02x\n", dense)
}