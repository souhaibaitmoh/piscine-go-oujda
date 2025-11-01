package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + (int(r) - '0')
	}
	return n
}

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		var start int

		if args[0] == "--upper" {
			start = 64
			args = args[1:]
		} else {
			start = 96
		}

		for i := range args {
			pos := atoi(args[i])
			if !(atoi(args[i]) < atoi("1") || atoi(args[i]) > atoi("26")) {
				c := pos + start
				z01.PrintRune(rune(c))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
