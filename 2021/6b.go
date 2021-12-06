package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/6.txt")
	if err != nil {
		log.Fatalf("Couldn't read: %v", err)
	}

	fish := make(map[int]int)
	for _, s := range strings.Split(string(f), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		fish[n]++
	}

	for i := 0; i < 256; i++ {
		newfish := make(map[int]int)
		for f, n := range fish {
			switch f {
			case 0:
				newfish[6] += n
				newfish[8] += n
			default:
				newfish[f-1] += n
			}
		}
		fish = newfish
	}

	sum := 0
	for _, n := range fish {
		sum += n
	}
	fmt.Println(sum)
}
