package piscine

func BasicAtoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		digit := int(ch - '0')
		res = res*10 + digit
	}
	return res
}
