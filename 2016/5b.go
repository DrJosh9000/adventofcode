package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

// Advent of Code 2016
// Day 5, part b

const chars = "0123456789abcdef"

func main() {
	doorID := []byte(os.Args[1])

	fmt.Println("!!! DECRYPTING !!!")
	password := []byte("________")
	fmt.Printf("%s", password)
	notdone := 0xff
	for i := 0; ; i++ {
		sum := md5.Sum(append(doorID, []byte(strconv.Itoa(i))...))
		if !(sum[0] == 0 && sum[1] == 0 && sum[2] < 8 && (notdone&(1<<sum[2]) != 0)) {
			continue
		}
		password[sum[2]] = chars[sum[3]>>4]
		notdone &^= 1 << sum[2]
		fmt.Printf("\r%s", password)
		if notdone == 0 {
			fmt.Println()
			return
		}
	}
}
