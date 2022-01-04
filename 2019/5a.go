package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/5.txt")
	if err != nil {
		log.Fatalf("Couldn't read input file: %v", err)
	}

	input := strings.Split(string(f), ",")
	m := make([]int, len(input))
	for i, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		m[i] = n
	}

	in := 1
	pc := 0
	pow10 := []int{1, 10, 100, 1000, 10000}
	opval := func(n int) int {
		mode := (m[pc] / pow10[n+1]) % 10
		switch mode {
		case 0: // position mode
			return m[m[pc+n]]
		case 1: // immediate mode
			return m[pc+n]
		}
		log.Fatalf("unimplemented mode")
		return 0
	}

vmLoop:
	for {
		switch m[pc] % 100 {
		case 1:
			m[m[pc+3]] = opval(1) + opval(2)
			pc += 4
		case 2:
			m[m[pc+3]] = opval(1) * opval(2)
			pc += 4
		case 3:
			m[m[pc+1]] = in
			pc += 2
		case 4:
			fmt.Println(opval(1))
			pc += 2
		case 99:
			break vmLoop
		default:
			log.Fatalf("Invalid opcode %d", m[pc])
		}
	}
}
