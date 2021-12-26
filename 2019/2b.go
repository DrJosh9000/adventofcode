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
	initial := make([]int, len(input))
	for i, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		initial[i] = n
	}

	for initial[1] = 0; initial[1] <= 99; initial[1]++ {
		for initial[2] = 0; initial[2] <= 99; initial[2]++ {
			m := make([]int, len(initial))
			copy(m, initial)
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
			if m[0] == 19690720 {
				fmt.Println(100*initial[1] + initial[2])
				return
			}
		}
	}
}
