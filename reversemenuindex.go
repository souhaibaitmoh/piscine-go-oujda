package piscine

func ReverseMenuIndex(menu []string) []string {
	s := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		s[i] = menu[len(menu)-1-i]
	}
	return s
}
