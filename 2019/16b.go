package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.ReadFile("inputs/16.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	inlen := len(f)

	offset, err := strconv.Atoi(string(f[:7]))
	if err != nil {
		log.Fatalf("First 7 digits somehow not a number: %v", err)
	}
	copies := (10000*inlen-offset)/inlen + 1
	offset %= inlen

	// Since the offset is well into the second half of the "real signal",
	// each new digit is just the (modulo-10) sum of the remaining digits...

	for i := range f {
		f[i] -= '0'
	}
	g := make([]byte, 0, copies*inlen)
	for i := 0; i < copies; i++ {
		g = append(g, f...)
	}
	g = g[offset:]

	for i := 0; i < 100; i++ {
		g = fakefft(g)
	}

	for i := 0; i < 8; i++ {
		g[i] += '0'
	}
	fmt.Println(string(g[:8]))
}

func fakefft(in []byte) []byte {
	out := make([]byte, len(in))
	var sum byte
	for i := range in {
		i := len(in) - i - 1
		sum += in[i]
		sum %= 10
		out[i] = sum
	}
	return out
}
