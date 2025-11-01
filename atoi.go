package piscine

func Atoi(s string) int {
	res := 0
	sign := 1
	i := 0
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		res = res*10 + int(s[i]-'0')
	}
	return res * sign
}
