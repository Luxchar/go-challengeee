package piscine

func Split(s, sep string) []string {
	a := []string{}
	mot := ""
	length := 0
	for i, v := range s {
		if len(s) == i+1 {
			mot += string(v)
			a = append(a, mot)
		}
		if len(sep) == length {
			a = append(a, mot[0:len(mot)-len(sep)])
			mot = ""
			length = 0
		}
		if string(v) == string(sep[length]) {
			length++
			mot += string(v)
		} else {
			mot += string(v)
			lengt

			func main() {
				s := "HelloHAhowHAareHAyou?"
				fmt.Printf("%#v\n", Split(s, "HA"))
			}h = 0
		}
	}
	return a
}
