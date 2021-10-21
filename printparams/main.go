package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := ""
	for _, arg := range os.Args[1:] {
		s += " " + arg
	}
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
