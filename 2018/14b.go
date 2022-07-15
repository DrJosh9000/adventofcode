package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

func main() {
	var N []int
	for _, c := range os.Args[1] {
		N = append(N, int(c - '0'))
	}
	
	elves := [2]int{0, 1}
	board := []int{3, 7}
	
	for {
		sum := board[elves[0]] + board[elves[1]]
		for _, c := range strconv.Itoa(sum) {
			board = append(board, int(c - '0'))
			if len(board) >= len(N) && slices.Equal(N, board[len(board)-len(N):]) {
				fmt.Println(len(board) - len(N))
				return
			}
		}
		b := len(board)
		elves[0] += board[elves[0]] + 1
		elves[0] %= b
		elves[1] += board[elves[1]] + 1
		elves[1] %= b
	}
}