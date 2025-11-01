package piscine

func Max(a []int) int {
	if len(a) > 0 {
		max := a[0]
		for _, k := range a {
			if k > max {
				max = k
			}
		}
		return max
	}
	return 0
}
