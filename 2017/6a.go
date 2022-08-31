package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

const numBanks = 16

type config [numBanks]int

func main() {
	var banks config
	exp.MustForEachLineIn("inputs/6.txt", func(line string) {
		for i, ns := range strings.Fields(line) {
			n, err := strconv.Atoi(ns)
			if err != nil {
				log.Fatalf("Couldn't atoi: %v", err)
			}
			banks[i] = n
		}
	})
	//fmt.Println(banks)
	
	seen := algo.Set[config]{
		banks: {},
	}
	cycles := 0
	for {
		blocks, j := -1, -1
		for i, n := range banks {
			if n > blocks {
				blocks, j = n, i
			}
		}
		banks[j] = 0
		for blocks > 0 {
			j++
			j %= numBanks
			banks[j]++
			blocks--
		}
		cycles++
		
		//fmt.Println(banks)
		
		if _, s := seen[banks]; s {
			fmt.Println(cycles)
			return
		}
		seen.Insert(banks)
	}
}