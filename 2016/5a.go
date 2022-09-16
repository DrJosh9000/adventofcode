package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

// Advent of Code 2016
// Day 5, part a

const chars = "0123456789abcdef"

func main() {
	doorID := []byte(os.Args[1])
	var password []byte
	for i := 0; ; i++ {
		sum := md5.Sum(append(doorID, []byte(strconv.Itoa(i))...))
		if sum[0] == 0 && sum[1] == 0 && sum[2] < 0x10 {
			password = append(password, chars[sum[2]])
		}
		if len(password) == 8 {
			fmt.Println(string(password))
			return
		}
	}
}
