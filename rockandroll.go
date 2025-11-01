package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}
	divby2 := n%2 == 0
	divby3 := n%3 == 0

	switch {
	case divby2 && divby3:
		return "rock and roll\n"
	case divby2:
		return "rock\n"
	case divby3:
		return "roll\n"
	default:
		return "error: non divisible\n"
	}
}
