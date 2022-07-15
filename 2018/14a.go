package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Usage: 14a number. Error was: %v", err)
	}
	
	elves := [2]int{0, 1}
	board := []int{3, 7}
	
	for len(board) < N + 10 {
		sum := board[elves[0]] + board[elves[1]]
		for _, c := range strconv.Itoa(sum) {
			board = append(board, int(c - '0'))
		}
		b := len(board)
		elves[0] += board[elves[0]] + 1
		elves[0] %= b
		elves[1] += board[elves[1]] + 1
		elves[1] %= b
	}
	
	for _, n := range board[N:N+10] {
		fmt.Print(n)
	}
	fmt.Println()
}