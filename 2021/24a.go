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
			registers[token[1]] = doAdd(op1, op2)
		case "mul":
			registers[token[1]] = doMul(op1, op2)
		case "div":
			registers[token[1]] = doDiv(op1, op2)
		case "mod":
			registers[token[1]] = doMod(op1, op2)
		case "eql":
			registers[token[1]] = doEql(op1, op2)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	const mp = 32
	var wg sync.WaitGroup
	wg.Add(mp)
	z := registers["z"]
	chunk := pow9[14] / mp
	for i := 0; i < mp; i++ {
		i := i
		lo, hi := chunk*i, chunk*(i+1)
		go func() {
			var best int
			for m := hi; m >= lo; m-- {
				n, err := z.eval(m)
				if err != nil {
					continue
				}
				if n == 0 {
					fmt.Println("goroutine:", i, "best:", best)
					wg.Done()
					return
				}
			}
		}()
	}
	wg.Wait()
}

var pow9 = make([]int, 15)

func init() {
	x := 1
	for i := range pow9 {
		pow9[i] = x
		x *= 9
	}
}

type expr interface {
	eval(int) (int, error)
}

type variable byte // variable representing digit 0-13 of the input

func (v variable) eval(m int) (int, error) {
	return (int(m)/pow9[13-v])%9 + 1, nil
}

type constant int // some literal in the program

func (c constant) eval(int) (int, error) { return int(c), nil }

func doAdd(a, b expr) expr {
	if a == constant(0) {
		return b
	}
	if b == constant(0) {
		return a
	}
	ac, aok := a.(constant)
	bc, bok := b.(constant)
	if aok && bok {
		return ac + bc
	}
	return add{a, b}
}

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

func doMul(a, b expr) expr {
	if a == constant(0) || b == constant(0) {
		return constant(0)
	}
	if a == constant(1) {
		return b
	}
	if b == constant(1) {
		return a
	}
	ac, aok := a.(constant)
	bc, bok := b.(constant)
	if aok && bok {
		return ac * bc
	}
	return mul{a, b}
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

func doDiv(a, b expr) expr {
	if b == constant(1) {
		return a
	}
	ac, aok := a.(constant)
	bc, bok := b.(constant)
	if aok && bok {
		if bc == 0 {
			log.Fatal("Input program always divides by zero")
		}
		return ac / bc
	}
	return div{a, b}
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

func doMod(a, b expr) expr {
	ac, aok := a.(constant)
	bc, bok := b.(constant)
	if aok && bok {
		if ac < 0 || bc <= 0 {
			log.Fatal("Input program always performs invalid modulus")
		}
		return ac % bc
	}
	return div{a, b}
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

func doEql(a, b expr) expr {
	ac, aok := a.(constant)
	bc, bok := b.(constant)
	if aok && bok {
		if ac == bc {
			return constant(1)
		}
		return constant(0)
	}
	return eql{a, b}
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
