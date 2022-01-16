package main

import (
	"bytes"
	"fmt"
	"log"
	"math/bits"
	"os"
)

func main() {
	input, err := os.ReadFile("inputs/24.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	x := parse(input)
	seen := map[state]bool{x: true}
	for {
		x = x.evolve()
		//fmt.Println(x)
		if seen[x] {
			fmt.Println(x)
			fmt.Println(int(x))
			break
		}
		seen[x] = true
	}
}

var neigh = []state{
	0: 0b00000_00000_00000_00001_00010,
	1: 0b00000_00000_00000_00010_00101,
	2: 0b00000_00000_00000_00100_01010,
	3: 0b00000_00000_00000_01000_10100,
	4: 0b00000_00000_00000_10000_01000,

	5: 0b00000_00000_00001_00010_00001,
	6: 0b00000_00000_00010_00101_00010,
	7: 0b00000_00000_00100_01010_00100,
	8: 0b00000_00000_01000_10100_01000,
	9: 0b00000_00000_10000_01000_10000,

	10: 0b00000_00001_00010_00001_00000,
	11: 0b00000_00010_00101_00010_00000,
	12: 0b00000_00100_01010_00100_00000,
	13: 0b00000_01000_10100_01000_00000,
	14: 0b00000_10000_01000_10000_00000,

	15: 0b00001_00010_00001_00000_00000,
	16: 0b00010_00101_00010_00000_00000,
	17: 0b00100_01010_00100_00000_00000,
	18: 0b01000_10100_01000_00000_00000,
	19: 0b10000_01000_10000_00000_00000,

	20: 0b00010_00001_00000_00000_00000,
	21: 0b00101_00010_00000_00000_00000,
	22: 0b01010_00100_00000_00000_00000,
	23: 0b10100_01000_00000_00000_00000,
	24: 0b01000_10000_00000_00000_00000,
}

type state int

func (s state) evolve() state {
	var t state
	for i := range neigh {
		p := bits.OnesCount(uint(s & neigh[i]))
		if p == 1 || (s&(1<<i) == 0 && p == 2) {
			t |= 1 << i
		}
	}
	return t
}

func parse(in []byte) state {
	var x state
	for j, row := range bytes.Split(in, []byte("\n")) {
		for i, b := range row {
			if b == '#' {
				x |= (1 << (i + 5*j))
			}
		}
	}
	return x
}

func (s state) String() string {
	b := []byte(".....\n.....\n.....\n.....\n.....")
	for i := 0; i < 25; i++ {
		if s&(1<<i) != 0 {
			b[i+i/5] = '#'
		}
	}
	return string(b)
}
