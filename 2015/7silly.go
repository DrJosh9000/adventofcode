package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2015
// Day 7, part a

// The silliest way possible?

type topic[T any] struct {
	mu     sync.Mutex
	subs   []chan T
	closed bool
}

func (q *topic[T]) Pub(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		panic("topic is closed")
	}
	for _, c := range q.subs {
		c <- v
	}
}

func (q *topic[T]) Sub() <-chan T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		panic("topic is closed")
	}
	ch := make(chan T, 1)
	q.subs = append(q.subs, ch)
	return ch
}

func (q *topic[T]) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, c := range q.subs {
		close(c)
	}
	q.subs = nil
	q.closed = true
}

var wires = make(map[string]*topic[uint16])

func wire(t string) *topic[uint16] {
	if _, w := wires[t]; !w {
		wires[t] = new(topic[uint16])
	}
	return wires[t]
}

func expr(t string) any {
	n, err := strconv.ParseUint(t, 10, 16)
	if err != nil {
		return wire(t).Sub()
	}
	return uint16(n)
}

func unary(start <-chan struct{}, x any, out *topic[uint16], op func(uint16) uint16) {
	switch x := x.(type) {
	case uint16:
		go func() {
			<-start
			out.Pub(op(x))
			out.Close()
		}()
	case <-chan uint16:
		go func() {
			<-start
			lo := op(<-x)
			out.Pub(lo)
			for {
				t, open := <-x
				if !open {
					break
				}
				if o := op(t); o != lo {
					out.Pub(o)
					lo = o
				}
			}
			out.Close()
		}()
	}
}

func curryL[T any](x T, op func(T, T) T) func(T) T {
	return func(t T) T { return op(x, t) }
}

func curryR[T any](y T, op func(T, T) T) func(T) T {
	return func(t T) T { return op(t, y) }
}

func binary(start <-chan struct{}, x, y any, out *topic[uint16], op func(uint16, uint16) uint16) {
	if x, ok := x.(uint16); ok {
		unary(start, y, out, curryL(x, op))
		return
	}
	if y, ok := y.(uint16); ok {
		unary(start, x, out, curryR(y, op))
		return
	}
	xch, ych := x.(<-chan uint16), y.(<-chan uint16)
	go func() {
		<-start
		lx, ly := <-xch, <-ych
		lo := op(lx, ly)
		out.Pub(lo)
		for {
			select {
			case x, open := <-xch:
				if open {
					lx = x
				} else {
					xch = nil
				}
			case y, open := <-ych:
				if open {
					ly = y
				} else {
					ych = nil
				}
			}
			if xch == nil && ych == nil {
				break
			}
			if o := op(lx, ly); o != lo {
				out.Pub(o)
				lo = o
			}
		}
		out.Close()
	}()
}

func id(x uint16) uint16     { return x }
func not(x uint16) uint16    { return ^x }
func and(x, y uint16) uint16 { return x & y }
func or(x, y uint16) uint16  { return x | y }
func lsh(x, y uint16) uint16 { return x << y }
func rsh(x, y uint16) uint16 { return x >> y }

func main() {
	start := make(chan struct{})
	for _, line := range exp.MustReadLines("inputs/7.txt") {
		tokens := strings.Fields(line)
		switch {
		case tokens[1] == "->":
			unary(start, expr(tokens[0]), wire(tokens[2]), id)
		case tokens[0] == "NOT":
			unary(start, expr(tokens[1]), wire(tokens[3]), not)
		case tokens[1] == "AND":
			binary(start, expr(tokens[0]), expr(tokens[2]), wire(tokens[4]), and)
		case tokens[1] == "OR":
			binary(start, expr(tokens[0]), expr(tokens[2]), wire(tokens[4]), or)
		case tokens[1] == "LSHIFT":
			binary(start, expr(tokens[0]), expr(tokens[2]), wire(tokens[4]), lsh)
		case tokens[1] == "RSHIFT":
			binary(start, expr(tokens[0]), expr(tokens[2]), wire(tokens[4]), rsh)
		default:
			log.Fatalf("Unknown form %q", line)
		}
	}

	ach := wire("a").Sub()
	close(start)

	fmt.Println(<-ach)
}
