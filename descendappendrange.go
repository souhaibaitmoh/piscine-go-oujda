package piscine

func DescendAppendRange(max, min int) []int {
	s := []int{}
	if min >= max {
		return []int{}
	}
	for i := max; i > min; i-- {
		s = append(s, i)
	}
	return s
}
