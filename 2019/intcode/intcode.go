package intcode

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type VM struct {
	M       []int
	In, Out chan int
}

func (vm *VM) Run() {
	m := make([]int, len(vm.M))
	copy(m, vm.M)

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
			m[m[pc+1]] = <-vm.In
			pc += 2
		case 4:
			vm.Out <- opval(1)
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
	close(vm.Out)
}

func ReadProgram(path string) []int {
	f, err := os.ReadFile(path)
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
	return m
}
