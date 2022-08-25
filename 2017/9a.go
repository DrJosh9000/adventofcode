package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	stream, err := os.ReadFile("inputs/9.txt")
	if err != nil {
		log.Fatalf("Couldn't read file: %v", err)
	}

	score := 0
	var stack []byte
	for _, c := range stream {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		switch stack[len(stack)-1] {
		case '{':
			switch c {
			case '{', '<', '!':
				stack = append(stack, c)
			case '}':
				score += len(stack)
				stack = stack[:len(stack)-1]
			}
		case '<':
			switch c {
			case '!':
				stack = append(stack, c)
			case '>':
				stack = stack[:len(stack)-1]
			}
		case '!':
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println(score)
}