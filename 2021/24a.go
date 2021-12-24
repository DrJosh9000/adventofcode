package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	f, err := os.Open("inputs/24.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	registers := map[string]expr{
		"w": constant(0),
		"x": constant(0),
		"y": constant(0),
		"z": constant(0),
	}
	var nextvar variable
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		token := strings.Split(sc.Text(), " ")
		if token[0] == "inp" {
			registers[token[1]] = nextvar
			nextvar++
			continue
		}
		op1 := registers[token[1]]
		if op1 == nil {
			log.Fatalf("first operand must be a register, got %q", token[1])
		}
		op2 := registers[token[2]]
		if op2 == nil {
			n, err := strconv.Atoi(token[2])
			if err != nil {
				log.Fatalf("second operand is neither a register nor a literal: %v", err)
			}
			op2 = constant(n)
		}
		switch token[0] {
		case "add":
			registers[token[1]] = add{op1, op2}
		case "mul":
			registers[token[1]] = mul{op1, op2}
		case "div":
			registers[token[1]] = div{op1, op2}
		case "mod":
			registers[token[1]] = mod{op1, op2}
		case "eql":
			registers[token[1]] = eql{op1, op2}
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	const mp = 32
	var wg sync.WaitGroup
	wg.Add(mp)
	z := registers["z"]
	chunk := pow10[14] / mp
	for i := 0; i < mp; i++ {
		i := i
		lo, hi := chunk*i, chunk*(i+1)
		go func() {
			var best int
			for m := lo; m < hi; m++ {
				n, err := z.eval(m)
				if err != nil {
					continue
				}
				if n == 0 {
					// m only ever increases...
					best = m
				}
			}
			fmt.Println("goroutine:", i, "best:", best)
			wg.Done()
		}()
	}
	wg.Wait()
}

var pow10 = make([]int, 15)

func init() {
	x := 1
	for i := range pow10 {
		pow10[i] = x
		x *= 10
	}
}

type expr interface {
	eval(int) (int, error)
}

type variable byte // variable representing digit 0-13 of the input

var errInvalidInput = errors.New("invalid input")

func (v variable) eval(m int) (int, error) {
	d := (int(m) / pow10[13-v]) % 10
	if d == 0 {
		return 0, errInvalidInput
	}
	return d, nil
}

type constant int // some literal in the program

func (c constant) eval(int) (int, error) { return int(c), nil }

type add struct {
	x, y expr
}

func (o add) eval(m int) (int, error) {
	a, err := o.x.eval(m)
	if err != nil {
		return 0, err
	}
	b, err := o.y.eval(m)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}

type mul struct {
	x, y expr
}

func (o mul) eval(m int) (int, error) {
	a, err := o.x.eval(m)
	if err != nil {
		return 0, err
	}
	b, err := o.y.eval(m)
	if err != nil {
		return 0, err
	}
	return a * b, nil
}

type div struct {
	x, y expr
}

var errDivisionByZero = errors.New("division by zero")

func (o div) eval(m int) (int, error) {
	a, err := o.x.eval(m)
	if err != nil {
		return 0, err
	}
	b, err := o.y.eval(m)
	if err != nil {
		return 0, err
	}
	if b == 0 {
		return 0, errDivisionByZero
	}
	return a / b, nil
}

type mod struct {
	x, y expr
}

var errInvalidMod = errors.New("invalid mod operation")

func (o mod) eval(m int) (int, error) {
	a, err := o.x.eval(m)
	if err != nil {
		return 0, err
	}
	b, err := o.y.eval(m)
	if err != nil {
		return 0, err
	}
	if a < 0 || b <= 0 {
		return 0, errInvalidMod
	}
	return a % b, nil
}

type eql struct {
	x, y expr
}

func (e eql) eval(m int) (int, error) {
	a, err := e.x.eval(m)
	if err != nil {
		return 0, err
	}
	b, err := e.y.eval(m)
	if err != nil {
		return 0, err
	}
	if a == b {
		return 1, nil
	}
	return 0, nil
}
