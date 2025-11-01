package piscine

func StringToIntSlice(str string) []int {
	res := []int(nil)
	for _, v := range str {
		res = append(res, int(v))
	}
	return res
}
