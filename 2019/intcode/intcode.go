package intcode

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type VM struct {
	m      map[int]int
	pc, rb int
}

var pow10 = []int{1, 10, 100, 1000, 10000}

func (vm *VM) Copy() *VM {
	m := make(map[int]int, len(vm.m))
	for i, x := range vm.m {
		m[i] = x
	}
	return &VM{m: m, pc: vm.pc, rb: vm.rb}
}

func (vm *VM) Peek(a int) int { return vm.m[a] }

func (vm *VM) Poke(a, v int) { vm.m[a] = v }

func (vm *VM) Run(in <-chan int, out chan<- int) {
	defer close(out)

	opaddr := func(n int) int {
		mode := (vm.m[vm.pc] / pow10[n+1]) % 10
		switch mode {
		case 0: // position mode
			return vm.m[vm.pc+n]
		case 1: // immediate mode
			return vm.pc + n
		case 2: // relative mode
			return vm.rb + vm.m[vm.pc+n]
		}
		log.Fatalf("Unimplemented mode %d", mode)
		return 0
	}

	for {
		switch vm.m[vm.pc] % 100 {
		case 1:
			vm.m[opaddr(3)] = vm.m[opaddr(1)] + vm.m[opaddr(2)]
			vm.pc += 4
		case 2:
			vm.m[opaddr(3)] = vm.m[opaddr(1)] * vm.m[opaddr(2)]
			vm.pc += 4
		case 3:
			t, ok := <-in
			if !ok {
				// Input channel closed. Returning temporarily halts the machine
				// in its current state so it could be resumed with new input
				// and output channels later on.
				return
			}
			vm.m[opaddr(1)] = t
			vm.pc += 2
		case 4:
			out <- vm.m[opaddr(1)]
			vm.pc += 2
		case 5:
			if vm.m[opaddr(1)] != 0 {
				vm.pc = vm.m[opaddr(2)]
			} else {
				vm.pc += 3
			}
		case 6:
			if vm.m[opaddr(1)] == 0 {
				vm.pc = vm.m[opaddr(2)]
			} else {
				vm.pc += 3
			}
		case 7:
			if vm.m[opaddr(1)] < vm.m[opaddr(2)] {
				vm.m[opaddr(3)] = 1
			} else {
				vm.m[opaddr(3)] = 0
			}
			vm.pc += 4
		case 8:
			if vm.m[opaddr(1)] == vm.m[opaddr(2)] {
				vm.m[opaddr(3)] = 1
			} else {
				vm.m[opaddr(3)] = 0
			}
			vm.pc += 4
		case 9: // "relative base offset"
			vm.rb += vm.m[opaddr(1)]
			vm.pc += 2
		case 99:
			return
		default:
			log.Fatalf("Invalid opcode %d", vm.m[vm.pc])
		}
	}
}

func ReadProgram(path string) *VM {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't read input file: %v", err)
	}

	input := strings.Split(string(f), ",")
	m := make(map[int]int, len(input))
	for i, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		m[i] = n
	}
	return &VM{m: m}
}
