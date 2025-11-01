package piscine

import "github.com/01-edu/z01"

func Putnbr(nb int) {
	if nb < 10 {
		z01.PrintRune('0')
		z01.PrintRune(rune(nb + '0'))
	} else {
		z01.PrintRune(rune(nb/10 + '0'))
		z01.PrintRune(rune(nb%10 + '0'))
	}
}

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			Putnbr(i)
			z01.PrintRune(' ')
			Putnbr(j)
			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
