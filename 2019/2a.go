package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/2.txt")
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

	m[1] = 12
	m[2] = 2

	pc := 0
vmLoop:
	for {
		switch m[pc] {
		case 1:
			m[m[pc+3]] = m[m[pc+1]] + m[m[pc+2]]
		case 2:
			m[m[pc+3]] = m[m[pc+1]] * m[m[pc+2]]
		case 99:
			break vmLoop
		default:
			log.Fatalf("Invalid opcode %d", m[pc])
		}
		pc += 4
	}

	fmt.Println(m[0])
}
