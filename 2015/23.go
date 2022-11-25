package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/emu"
)

// Advent of Code 2015
// Day 23, part a

func main() {
	program := exp.MustReadLines("inputs/23.txt")

	run, err := emu.Transpile(program, map[string]emu.TranslatorFunc{
		"hlf": func(line int, args []string) (impl string, jumpTargets []int, err error) {
			return fmt.Sprintf("r[%d] /= 2", args[0][0]-'a'), nil, nil
		},
		"tpl": func(line int, args []string) (impl string, jumpTargets []int, err error) {
			return fmt.Sprintf("r[%d] *= 3", args[0][0]-'a'), nil, nil
		},
		"inc": func(line int, args []string) (impl string, jumpTargets []int, err error) {
			return fmt.Sprintf("r[%d]++", args[0][0]-'a'), nil, nil
		},
		"jmp": func(line int, args []string) (impl string, jumpTargets []int, err error) {
			o := exp.Must(strconv.Atoi(args[0]))
			return fmt.Sprintf("goto l%d", o+line), []int{o + line}, nil
		},
		"jie": func(line int, args []string) (impl string, jumpTargets []int, err error) {
			o := exp.Must(strconv.Atoi(args[1]))
			return fmt.Sprintf("if r[%d] %% 2 == 0 { goto l%d }", args[0][0]-'a', o+line), []int{o + line}, nil
		},
		"jio": func(line int, args []string) (impl string, jumpTargets []int, err error) {
			o := exp.Must(strconv.Atoi(args[1]))
			return fmt.Sprintf("if r[%d] == 1 { goto l%d }", args[0][0]-'a', o+line), []int{o + line}, nil
		},
	})

	if err != nil {
		log.Fatalf("Couldn't transpile: %v", err)
	}

	r := make([]int, 2)
	run(r, nil, nil, nil)
	fmt.Println(r[1])

	r[0], r[1] = 1, 0
	run(r, nil, nil, nil)
	fmt.Println(r[1])
}
