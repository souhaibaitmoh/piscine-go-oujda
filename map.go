package piscine

func Map(f func(int) bool, a []int) (res []bool) {
	for _, v := range a {
		res = append(res, f(v))
	}
	return
}
