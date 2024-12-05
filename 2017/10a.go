package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

func main() {
	// Getting func-y... doing it this way for fun(c)... not for readability or clarity.
	// Read the file, cast the contents to a string, split by comma,
	// trim spaces from each item, and finally convert each item from a decimal string to an int.
	lengths := algo.Map(
		algo.Map(
			strings.Split(
				string(exp.Must(os.ReadFile("inputs/10.txt"))),
				",",
			), 
			strings.TrimSpace,
		), 
		exp.MustFunc(strconv.Atoi),
	)
	
	circle := make([]int, 256)
	for i := range circle {
		circle[i] = i
	}
	pos, skip := 0, 0
		
	for _, l := range lengths {
		for i := 0; i < l/2; i++ {
			j, k := (pos+i)%len(circle), (pos+l-i-1)%len(circle)
			circle[j], circle[k] = circle[k], circle[j]
		}
		pos += l + skip
		skip++
	}
	
	fmt.Println(circle[0] * circle[1])
}