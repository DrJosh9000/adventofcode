package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in, err := os.ReadFile("inputs/5.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	
	var reps []string
	for c := 'A'; c <= 'Z'; c++ {
		reps = append(reps, 
			fmt.Sprintf("%c%c", c, c+32), "",
			fmt.Sprintf("%c%c", c+32, c), "",
		)
	}
	
	r := strings.NewReplacer(reps...)
	s := strings.TrimSpace(string(in))
	i := 0
	for  {
		//fmt.Println(len(s))
		n := r.Replace(s)
		if n == s {
			break
		}
		s = n
		i++
	}
	fmt.Println(i, "iterations")
	fmt.Println(len(s))
}