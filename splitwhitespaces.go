package piscine

func SplitWhiteSpaces(s string) []string {
	a := []string{}
	mot := ""
	index := 1
	for _, i := range s {
		if len(s) == index-1 {
			mot += string(i)
			a = append(a, mot)
		}
		if string(i) == " " {
			if mot != "" {
				a = append(a, mot)
				mot = ""
			}
			mot = ""
		} else {
			mot += string(i)
		}
		index++
	}
	return a
}
