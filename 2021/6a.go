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

	var fish []int
	for _, s := range strings.Split(string(f), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't atoi: %v", err)
		}
		fish = append(fish, n)
	}

	for i := 0; i < 80; i++ {
		var newfish []int
		for _, f := range fish {
			switch f {
			case 0:
				newfish = append(newfish, 6, 8)
			default:
				newfish = append(newfish, f-1)
			}
		}
		fish = newfish
	}

	fmt.Println(len(fish))
}
