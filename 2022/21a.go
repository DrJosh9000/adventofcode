package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 21, part a

// Let's copy nearly all of 2015/7silly.go...

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

var wires = make(map[string]*topic[int])

func wire(t string) *topic[int] {
	if _, w := wires[t]; !w {
		wires[t] = new(topic[int])
	}
	return wires[t]
}

func expr(t string) any {
	n, err := strconv.Atoi(t)
	if err != nil {
		return wire(t).Sub()
	}
	return n
}

func unary(start <-chan struct{}, x any, out *topic[int], op func(int) int) {
	switch x := x.(type) {
	case int:
		go func() {
			<-start
			out.Pub(op(x))
			out.Close()
		}()
	case <-chan int:
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

func binary(start <-chan struct{}, x, y any, out *topic[int], op func(int, int) int) {
	if x, ok := x.(int); ok {
		unary(start, y, out, curryL(x, op))
		return
	}
	if y, ok := y.(int); ok {
		unary(start, x, out, curryR(y, op))
		return
	}
	xch, ych := x.(<-chan int), y.(<-chan int)
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

var ops = map[string]func(int, int) int{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

func main() {
	start := make(chan struct{})
	for _, line := range exp.MustReadLines("inputs/21.txt") {
		name, op, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("Couldn't find cut string in %q", line)
		}
		out := wire(name)
		n, err := strconv.Atoi(op)
		if err == nil {
			go func() {
				<-start
				out.Pub(n)
				out.Close()
			}()
			continue
		}

		opf := strings.Fields(op)
		binary(start, expr(opf[0]), expr(opf[2]), out, ops[opf[1]])
	}
	rootCh := wire("root").Sub()
	close(start)

	fmt.Println(<-rootCh)
}
