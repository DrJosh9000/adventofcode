package main

import (
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2023
// Day 15, part b

const inputPath = "2023/inputs/15.txt"

func main() {
	lines := exp.MustReadLines(inputPath)

	type lens struct {
		fl    int
		label string
	}
	boxes := make([][]lens, 256)
	for _, line := range lines {
	tokenLoop:
		for _, token := range strings.Split(line, ",") {
			if label, ok := strings.CutSuffix(token, "-"); ok {
				idx := hash(label)
				box := boxes[idx]
				for i, l := range box {
					if l.label == label {
						boxes[idx] = append(box[:i], box[i+1:]...)
						break
					}
				}
				continue
			}
			label, fls, ok := strings.Cut(token, "=")
			if !ok {
				panic("string has neither - nor =")
			}
			fl := exp.MustAtoi(fls)
			idx := hash(label)
			box := boxes[idx]
			for i, l := range box {
				if l.label == label {
					box[i].fl = fl
					continue tokenLoop
				}
			}
			boxes[idx] = append(box, lens{fl: fl, label: label})
		}
	}

	sum := 0
	for b, box := range boxes {
		for s, lens := range box {
			sum += (1 + b) * (1 + s) * lens.fl
		}
	}
	fmt.Println(sum)
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h += int(c)
		h *= 17
		h %= 256
	}
	return h
}
