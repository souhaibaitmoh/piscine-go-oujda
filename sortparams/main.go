package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args[1:]

	for range ar {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}

	for _, arg := range ar {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
