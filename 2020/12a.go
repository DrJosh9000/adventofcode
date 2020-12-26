package main

import (
	"fmt"
	"io"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func abs(x int) int {
	if x<0 {
		return -x
	}
	return x
}

func main() {
	f, err := os.Open("input.12")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	type ins struct {
		opcode rune
		operand int
	}
	var x, y, rot int
	for {
		in := &ins{}
		_, err := fmt.Fscanf(f, "%c%d\n", &in.opcode, &in.operand)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		switch in.opcode {
		case 'N':
			y+=in.operand
		case 'S':
			y-=in.operand
		case 'E':
			x+=in.operand
		case 'W':
			x-=in.operand
		case 'L':
			rot+=in.operand
		case 'R':
			rot-=in.operand
		case 'F':
			for rot<0 {rot+=360}
			for rot>=360 {rot-=360}
			switch rot {
			case 0:
				x+=in.operand
			case 90:
				y+=in.operand
			case 180:
				x-=in.operand
			case 270:
				y-=in.operand
			}
		} 
	}
	fmt.Println(abs(x)+abs(y))
}
