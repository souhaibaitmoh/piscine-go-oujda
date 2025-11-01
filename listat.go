package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		count++
		l = l.Next
	}
	return nil
}
