package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("inputs/16.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	sig := make([]int, len(f))
	for i := range f {
		sig[i] = int(f[i] - '0')
	}

	for i := 0; i < 100; i++ {
		sig = fft(sig)
	}

	for i := range sig {
		f[i] = byte(sig[i] + '0')
	}
	fmt.Println(string(f[:8]))
}

var pattern = []int{0, 1, 0, -1}

func fft(in []int) []int {
	out := make([]int, len(in))
	for i := range out {
		for j, x := range in {
			out[i] += x * pattern[((j+1)/(i+1))%4]
		}
		if out[i] < 0 {
			out[i] = -out[i]
		}
		out[i] %= 10
	}
	return out
}
