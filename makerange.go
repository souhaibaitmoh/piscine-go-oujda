package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = min + i
	}
	return s
}
