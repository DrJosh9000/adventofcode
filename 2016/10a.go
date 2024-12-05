package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 10, part a

// An utterly gratuitous use of concurrency.

var bots = make(map[int]*bot)

type bot struct {
	input  chan int
	lt, ht string
	ln, hn int
}

func newBot() *bot {
	return &bot{
		input: make(chan int, 2),
	}
}

func (b *bot) run(n int) {
	for {
		x, y := <-b.input, <-b.input
		if x > y {
			x, y = y, x
		}
		if x == 17 && y == 61 {
			fmt.Println(n)
			os.Exit(0)
		}
		if b.lt == "bot" {
			bots[b.ln].input <- x
		}
		if b.ht == "bot" {
			bots[b.hn].input <- y
		}
	}
}

func main() {
	for _, line := range exp.MustReadLines("inputs/10.txt") {
		switch {
		case strings.HasPrefix(line, "value"):
			var v, b int
			exp.Must(fmt.Sscanf(line, "value %d goes to bot %d", &v, &b))
			if bots[b] == nil {
				bots[b] = newBot()
			}
			bots[b].input <- v
		case strings.HasPrefix(line, "bot"):
			var lt, ht string
			var b, ln, hn int
			exp.Must(fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &b, &lt, &ln, &ht, &hn))
			if bots[b] == nil {
				bots[b] = newBot()
			}
			bt := bots[b]
			bt.lt, bt.ln, bt.ht, bt.hn = lt, ln, ht, hn
		default:
			log.Fatalf("Don't understand %q", line)
		}
	}

	for n, bt := range bots {
		go bt.run(n)
	}
	select {}
}
