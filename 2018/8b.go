package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"os"
)

type node struct {
	children []node
	metadata []int
}

func parse(nums []int) (node, int) {
	var n node
	j := 2
	for i := 0; i < nums[0]; i++ {
		no, nj := parse(nums[j:])
		j += nj
		n.children = append(n.children, no)
	}
	end := j+nums[1]
	n.metadata = nums[j:end]
	return n, end
}

func (n node) sum() int {
	var s int
	if len(n.children) == 0 {
		for _, x := range n.metadata {
			s += x
		}
		return s
	}
	for _, x := range n.metadata {
		x--
		if x < 0 || x >= len(n.children) {
			continue
		}
		s += n.children[x].sum()
	}
	return s
}

func main() {
	in, err := os.ReadFile("inputs/8.txt")	
	if err != nil {
		log.Fatalf("Couldn't read file: %v", err)
	}
	
	var nums []int
	for _, sn := range strings.Split(string(in), " ") {
		n, err := strconv.Atoi(strings.TrimSpace(sn))
		if err != nil {
			log.Fatalf("Couldn't parse number: %v", err)
		}
		nums = append(nums, n)
	}
	
	root, _ := parse(nums)
	fmt.Println(root.sum())
}