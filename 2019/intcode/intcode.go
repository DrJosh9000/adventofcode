package intcode

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type VM []int

var pow10 = []int{1, 10, 100, 1000, 10000}

func (vm VM) Run(in <-chan int, out chan<- int) {
	m := append(VM{}, vm...)

	pc := 0
	opval := func(n int) int {
		mode := (m[pc] / pow10[n+1]) % 10
		switch mode {
		case 0: // position mode
			return m[m[pc+n]]
		case 1: // immediate mode
			return m[pc+n]
		}
		log.Fatalf("Unimplemented mode %d", mode)
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
			t, ok := <-in
			if !ok {
				log.Fatal("Input channel closed")
			}
			m[m[pc+1]] = t
			pc += 2
		case 4:
			out <- opval(1)
			pc += 2
		case 5:
			if opval(1) != 0 {
				pc = opval(2)
			} else {
				pc += 3
			}
		case 6:
			if opval(1) == 0 {
				pc = opval(2)
			} else {
				pc += 3
			}
		case 7:
			if opval(1) < opval(2) {
				m[m[pc+3]] = 1
			} else {
				m[m[pc+3]] = 0
			}
			pc += 4
		case 8:
			if opval(1) == opval(2) {
				m[m[pc+3]] = 1
			} else {
				m[m[pc+3]] = 0
			}
			pc += 4
		case 99:
			break vmLoop
		default:
			log.Fatalf("Invalid opcode %d", m[pc])
		}
	}
	close(out)
}

func ReadProgram(path string) VM {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't read input file: %v", err)
	}

	input := strings.Split(string(f), ",")
	m := make(VM, len(input))
	for i, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		m[i] = n
	}
	return m
}
