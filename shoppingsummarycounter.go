package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			m := str[start:i]
			result[m]++
			start = i + 1
		}
	}
	m := str[start:]
	result[m]++
	return result
}
