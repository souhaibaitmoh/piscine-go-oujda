package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	ns := []string{}

	for _, c := range s {
		if c != "" {
			ns = append(ns, c)
		}
	}
	*ptr = ns
	return len(ns)
}
