package piscine

func Join(strs []string, sep string) string {
	a := ""
	for _, v := range strs {
		a += v
		if v != strs[len(strs)] {
			a += sep
		}
	}
	return a
}
