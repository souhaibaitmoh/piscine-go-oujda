package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i, arg := range args {
		if i > 0 {
			z01.PrintRune('\n')
		}
		for _, r := range arg {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
